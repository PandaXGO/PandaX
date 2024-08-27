package nodes

import (
	"pandax/pkg/rule_engine/message"
)

type messageTypeSwitchNode struct {
	bareNode
}
type messageTypeSwitchNodeFactory struct{}

func (f messageTypeSwitchNodeFactory) Name() string     { return "MessageTypeSwitchNode" }
func (f messageTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f messageTypeSwitchNodeFactory) Labels() []string {
	return []string{
		message.RowMes,
		message.AttributesMes,
		message.TelemetryMes,
		message.RpcRequestFromDevice,
		message.RpcRequestToDevice,
		message.UpEventMes,
		message.ConnectMes,
		message.DisConnectMes,
	}
}
func (f messageTypeSwitchNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &messageTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *messageTypeSwitchNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")
	nodes := n.GetLinkedNodes()
	messageType := msg.MsgType
	for label, node := range nodes {
		if messageType == label {
			return node.Handle(msg)
		}
	}
	n.Debug(msg, message.DEBUGOUT, "消息类型不正确")
	return nil
}
