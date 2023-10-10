package nodes

import (
	"pandax/iothub/client/mqttclient"
	"pandax/pkg/rule_engine/message"
)

type rpcRequestNode struct {
	bareNode
	Timeout int                   `json:"timeout"`
	Payload mqttclient.RpcPayload `json:"payload"`
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

	var rpc = &mqttclient.RpcRequest{Client: mqttclient.MqttClient, Mode: "double", Timeout: n.Timeout}
	rpc.GetRequestId()
	err := rpc.RespondTpc(n.Payload)
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
