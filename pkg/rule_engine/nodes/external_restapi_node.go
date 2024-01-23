package nodes

import (
	"encoding/json"
	"errors"
	"pandax/kit/httpclient"
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
func (f externalRestapiNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &externalRestapiNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalRestapiNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")
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
			n.Debug(msg, message.DEBUGOUT, err.Error())
			return failureLableNode.Handle(msg)
		} else {
			if successLableNode != nil {
				metadata := msg.Metadata
				for key, value := range response {
					metadata.SetValue(key, value)
				}
				msg.Metadata = metadata
				n.Debug(msg, message.DEBUGOUT, "")
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
				n.Debug(msg, message.DEBUGOUT, "接口请求失败")
				return failureLableNode.Handle(msg)
			}
		} else {
			if successLableNode != nil {
				n.Debug(msg, message.DEBUGOUT, "")
				return successLableNode.Handle(msg)
			}
		}
	}
	return nil
}
