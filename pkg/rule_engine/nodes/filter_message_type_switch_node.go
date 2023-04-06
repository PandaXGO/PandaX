package nodes

import (
	"fmt"
	"pandax/pkg/rule_engine/message"

	"github.com/sirupsen/logrus"
)

type messageTypeSwitchNode struct {
	bareNode
}
type messageTypeSwitchNodeFactory struct{}

func (f messageTypeSwitchNodeFactory) Name() string     { return "MessageTypeSwitchNode" }
func (f messageTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f messageTypeSwitchNodeFactory) Labels() []string {
	return []string{
		message.MessageTypePostAttributesRequest,
		message.MessageTypePostTelemetryRequest,
		message.MessageTypeConnectEvent,
		message.MessageTypeDisconnectEvent,
		"Other",
	}
}
func (f messageTypeSwitchNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &messageTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *messageTypeSwitchNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	nodes := n.GetLinkedNodes()
	messageType := msg.GetType()

	for label, node := range nodes {
		if messageType == label {
			return node.Handle(msg)
		}
	}
	// 自定义类型 或 未识别类型
	if node := n.GetLinkedNode("Other"); node != nil {
		return node.Handle(msg)
	}
	// not found
	return fmt.Errorf("%s no label to handle message", n.Name())
}
