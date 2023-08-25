package nodes

import (
	"pandax/pkg/global"
	"pandax/pkg/mqtt"
	"pandax/pkg/rule_engine/message"
)

type rpcRespondNode struct {
	bareNode
	RequestId int `json:"requestId"`
}

type rpcRespondFactory struct{}

func (f rpcRespondFactory) Name() string     { return "RpcRespondNode" }
func (f rpcRespondFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f rpcRespondFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f rpcRespondFactory) Create(id string, meta Metadata) (Node, error) {
	node := &rpcRespondNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *rpcRespondNode) Handle(msg *message.Message) error {
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")
	RequestId := n.RequestId
	if RequestId == 0 {
		RequestId = int(msg.Metadata.GetValue("requestId").(float64))
	}
	var datas = mqtt.RpcPayload{
		Method: msg.Msg.GetValue("method").(string),
		Params: msg.Msg.GetValue("params"),
	}
	rpc := &mqtt.RpcRequest{Client: global.MqttClient, RequestId: RequestId}
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
