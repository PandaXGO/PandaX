package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type externalMqNode struct {
	bareNode
}

type externalMqNodeFactory struct{}

func (f externalMqNodeFactory) Name() string     { return "ExternalMqNode" }
func (f externalMqNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalMqNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalMqNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalMqNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalMqNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	return nil
}
