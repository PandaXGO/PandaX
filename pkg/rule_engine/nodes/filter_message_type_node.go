package nodes

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type messageTypeFilterNode struct {
	bareNode
	MessageTypes []string `json:"messageTypes" yaml:"messageTypes"`
}

type messageTypeFilterNodeFactory struct{}

func (f messageTypeFilterNodeFactory) Name() string     { return "MessageTypeFilterNode" }
func (f messageTypeFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f messageTypeFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &messageTypeFilterNode{
		bareNode:     newBareNode(f.Name(), id, meta, labels),
		MessageTypes: []string{},
	}
	return decodePath(meta, node)
}

func (n *messageTypeFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")
	if trueLabelNode == nil || falseLabelNode == nil {
		return fmt.Errorf("no true or false label linked node in %s", n.Name())
	}
	messageType := msg.GetType()

	// TODO: how to resolve user customized message type dynamically
	//userMessageType := msg.GetMetadata().GetKeyValue(n.Metadata().MessageTypeKey)
	userMessageType := "TODO"
	for _, filterType := range n.MessageTypes {
		if filterType == messageType || filterType == userMessageType {
			return trueLabelNode.Handle(msg)
		}
	}
	return falseLabelNode.Handle(msg)
}
