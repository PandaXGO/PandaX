package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type messageTypeFilterNode struct {
	bareNode
	MessageTypes []string `json:"messageTypes" yaml:"messageTypes"`
}

type messageTypeFilterNodeFactory struct{}

func (f messageTypeFilterNodeFactory) Name() string     { return "MessageTypeNode" }
func (f messageTypeFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f messageTypeFilterNodeFactory) Labels() []string { return []string{"True", "False"} }

func (f messageTypeFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &messageTypeFilterNode{
		bareNode:     newBareNode(f.Name(), id, meta, f.Labels()),
		MessageTypes: []string{},
	}
	return decodePath(meta, node)
}

func (n *messageTypeFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")
	messageType := msg.GetType()

	for _, filterType := range n.MessageTypes {
		if filterType == messageType && trueLabelNode != nil {
			return trueLabelNode.Handle(msg)
		}
	}
	if falseLabelNode != nil {
		return falseLabelNode.Handle(msg)
	}
	return nil
}
