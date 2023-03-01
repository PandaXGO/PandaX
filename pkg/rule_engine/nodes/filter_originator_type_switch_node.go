package nodes

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type originatorTypeSwitchNode struct {
	bareNode
}

type originatorTypeSwitchNodeFactory struct{}

func (f originatorTypeSwitchNodeFactory) Name() string     { return "OriginatorTypeSwitchNode" }
func (f originatorTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f originatorTypeSwitchNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{}
	node := &originatorTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *originatorTypeSwitchNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	nodes := n.GetLinkedNodes()
	originatorType := msg.GetOriginator()

	for label, node := range nodes {
		if originatorType == label {
			return node.Handle(msg)
		}
	}
	// not found
	return fmt.Errorf("%s no label to handle message", n.Name())
}
