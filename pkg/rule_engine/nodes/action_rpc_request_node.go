package nodes

import (
	"errors"
	"pandax/iothub/client/mqttclient"
	"pandax/pkg/rule_engine/message"
)

type rpcRequestNode struct {
	bareNode
	RequestId int `json:"requestId"`
}

type rpcRequestNodeFactory struct{}

func (f rpcRequestNodeFactory) Name() string     { return "RpcRequestNode" }
func (f rpcRequestNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f rpcRequestNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f rpcRequestNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &rpcRequestNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *rpcRequestNode) Handle(msg *message.Message) error {
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")
	RequestId := n.RequestId
	if RequestId == 0 {
		RequestId = int(msg.Metadata.GetValue("requestId").(float64))
	}
	if msg.Msg.GetValue("method") == nil || msg.Msg.GetValue("params") == nil {
		return errors.New("请求响应格式不正确")
	}
	var datas = mqttclient.RpcPayload{
		Method: msg.Msg.GetValue("method").(string),
		Params: msg.Msg.GetValue("params"),
	}
	var rpc = &mqttclient.RpcRequest{Client: mqttclient.MqttClient, Mode: "single"}
	rpc.GetRequestId()
	err := rpc.RespondTpc(datas)
	if err != nil {
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLableNode != nil {
		return successLableNode.Handle(msg)
	}
	return nil
}
