package nodes

import (
	"pandax/pkg/global"
	"pandax/pkg/mqtt"
	"pandax/pkg/rule_engine/message"
)

type rpcRequestNode struct {
	bareNode
	Timeout int             `json:"timeout"`
	Payload mqtt.RpcPayload `json:"payload"`
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

	var rpc = &mqtt.RpcRequest{Client: global.MqttClient, Mode: "double", Timeout: n.Timeout}
	rpc.GetRequestId()
	respPayload, err := rpc.RequestCmd(n.Payload)
	if err != nil {
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}
	msgM := msg.Msg
	msgM["payload"] = respPayload
	msg.Msg = msgM
	if successLableNode != nil {
		return successLableNode.Handle(msg)
	}
	return nil
}
