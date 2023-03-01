package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

const ScriptFilterNodeName = "ScriptFilterNode"

type scriptFilterNode struct {
	bareNode
	Scripts string `json:"scripts" yaml:"scripts"`
}

type scriptFilterNodeFactory struct{}

func (f scriptFilterNodeFactory) Name() string     { return "ScriptFilterNode" }
func (f scriptFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f scriptFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &scriptFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *scriptFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")
	scriptEngine := NewScriptEngine()
	isTrue, error := scriptEngine.ScriptOnFilter(msg, n.Scripts)
	if isTrue == true && error == nil {
		return trueLabelNode.Handle(msg)
	}
	return falseLabelNode.Handle(msg)
}
