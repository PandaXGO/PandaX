package nodes

import (
	"encoding/json"
	"errors"
	"github.com/XM-GO/PandaKit/httpclient"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type externalRestapiNode struct {
	bareNode
	RestEndpointUrlPattern string            `json:"restEndpointUrlPattern" yaml:"restEndpointUrlPattern"`
	RequestMethod          string            `json:"requestMethod" yaml:"requestMethod"`
	Headers                map[string]string `json:"headers" yaml:"headers"`
}

type externalRestapiNodeFactory struct{}

func (f externalRestapiNodeFactory) Name() string     { return "RestapiNode" }
func (f externalRestapiNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalRestapiNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalRestapiNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalRestapiNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalRestapiNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")
	if n.RequestMethod == "GET" {
		resp := httpclient.NewRequest(n.RestEndpointUrlPattern).Get()
		if resp.StatusCode != 200 {
			return errors.New("网络请求失败")
		}
		var response map[string]interface{}
		err := json.Unmarshal(resp.Body, &response)
		if err != nil && failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			if successLableNode != nil {
				metadata := msg.GetMetadata()
				for key, value := range response {
					metadata.SetKeyValue(key, value)
				}
				msg.SetMetadata(metadata)
				return successLableNode.Handle(msg)
			}
		}
	}
	if n.RequestMethod == "POST" {
		binary, _ := msg.MarshalBinary()
		req := httpclient.NewRequest(n.RestEndpointUrlPattern)
		for key, value := range n.Headers {
			req.Header(key, value)
		}
		resp := req.PostJson(string(binary))
		if resp.StatusCode != 200 {
			if failureLableNode != nil {
				return failureLableNode.Handle(msg)
			}
		} else {
			if successLableNode != nil {
				return successLableNode.Handle(msg)
			}
		}
	}
	/*if n.RequestMethod == "PUT" {
		binary, _ := msg.MarshalBinary()
		req := httpclient.NewRequest(n.RestEndpointUrlPattern)
		for key,value := range n.Headers {
			req.Header(key,value)
		}
		_, err := http.HttpPut(n.RestEndpointUrlPattern, n.Headers, nil, binary)
		if err != nil {
			if failureLableNode != nil {
				return failureLableNode.Handle(msg)
			}
		} else {
			if successLableNode != nil {
				return successLableNode.Handle(msg)
			}
		}
	}
	if n.RequestMethod == "DELETE" {
		_, err := http.HttpDelete(n.RestEndpointUrlPattern)
		if err != nil {
			if failureLableNode != nil {
				return failureLableNode.Handle(msg)
			}
		} else {
			if successLableNode != nil {
				return successLableNode.Handle(msg)
			}
		}
	}*/
	return nil
}
