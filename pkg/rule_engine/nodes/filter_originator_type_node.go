package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type originatorTypeFilterNode struct {
	bareNode
	Filters []string `json:"filters" yaml:"filters"`
}

type originatorFilterNodeFactory struct{}

func (f originatorFilterNodeFactory) Name() string     { return "OriginatorFilterNode" }
func (f originatorFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f originatorFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &originatorTypeFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
		Filters:  []string{},
	}
	return decodePath(meta, node)
}

func (n *originatorTypeFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")

	//links := n.GetLinks()
	originatorType := msg.GetOriginator()

	for _, filter := range n.Filters {
		if originatorType == filter {
			return trueLabelNode.Handle(msg)
		}
	}
	// not found
	return falseLabelNode.Handle(msg)
}
