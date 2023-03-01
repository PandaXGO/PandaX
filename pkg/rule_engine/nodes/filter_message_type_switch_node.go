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

func (f messageTypeSwitchNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &messageTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *messageTypeSwitchNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	nodes := n.GetLinkedNodes()
	messageType := msg.GetType()
	messageTypeKey, _ := n.Metadata().Value(NODE_CONFIG_MESSAGE_TYPE_KEY)
	userMessageType := msg.GetMetadata().GetKeyValue(messageTypeKey.(string))

	for label, node := range nodes {
		if messageType == label || userMessageType == label {
			return node.Handle(msg)
		}
	}
	// if other label exist
	if node := n.GetLinkedNode("Other"); node != nil {
		return node.Handle(msg)
	}

	// not found
	return fmt.Errorf("%s no label to handle message", n.Name())
}
