package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/flow_engine/message"
)

type externalSendSmsNode struct {
	bareNode
}

type externalSendSmsNodeFactory struct{}

func (f externalSendSmsNodeFactory) Name() string     { return "ExternalSendSmslNode" }
func (f externalSendSmsNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalSendSmsNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalSendSmsNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalSendSmsNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalSendSmsNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	//failureLabelNode := n.GetLinkedNode("Failure")

	return successLabelNode.Handle(msg)
}
