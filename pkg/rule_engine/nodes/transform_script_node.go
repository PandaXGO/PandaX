package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type transformScriptNode struct {
	bareNode
	Script string `json:"script" yaml:"script"`
}

type transformScriptNodeFactory struct{}

func (f transformScriptNodeFactory) Name() string     { return "ScriptNode" }
func (f transformScriptNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }
func (f transformScriptNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f transformScriptNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}
	node := &transformScriptNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *transformScriptNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")

	scriptEngine := NewScriptEngine()
	newMessage, err := scriptEngine.ScriptOnMessage(msg, n.Script)
	if err != nil {
		return failureLabelNode.Handle(msg)
	}
	return successLabelNode.Handle(newMessage)
}
