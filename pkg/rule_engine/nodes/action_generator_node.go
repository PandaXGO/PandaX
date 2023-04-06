package nodes

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type messageGeneratorNode struct {
	bareNode
	DetailBuilderScript string `json:"detail_builder_script" yaml:"detail_builder_script"`
	FrequenceInSecond   int32  `json:"frequency" yaml:"frequency"`
}

type messageGeneratorNodeFactory struct{}

func (f messageGeneratorNodeFactory) Name() string     { return "MessageGeneratorNode" }
func (f messageGeneratorNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f messageGeneratorNodeFactory) Labels() []string { return []string{"Created", "Updated"} }
func (f messageGeneratorNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &messageGeneratorNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *messageGeneratorNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	createdLabelNode := n.GetLinkedNode("Created")
	updatedLabelNode := n.GetLinkedNode("Updated")
	if createdLabelNode == nil || updatedLabelNode == nil {
		return fmt.Errorf("no valid label linked node in %s", n.Name())
	}

	return nil
}
