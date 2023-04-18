package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

const ScriptFilterNodeName = "ScriptFilterNode"

type scriptFilterNode struct {
	bareNode
	Script string `json:"script" yaml:"script"`
}

type scriptFilterNodeFactory struct{}

func (f scriptFilterNodeFactory) Name() string     { return ScriptFilterNodeName }
func (f scriptFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f scriptFilterNodeFactory) Labels() []string { return []string{"True", "False"} }
func (f scriptFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &scriptFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *scriptFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")
	scriptEngine := NewScriptEngine(msg, "Filter", n.Script)
	isTrue, error := scriptEngine.ScriptOnFilter()
	if isTrue == true && error == nil && trueLabelNode != nil {
		return trueLabelNode.Handle(msg)
	} else {
		if falseLabelNode != nil {
			return falseLabelNode.Handle(msg)
		}
	}

	return nil
}
