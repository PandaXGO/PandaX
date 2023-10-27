package nodes

import (
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
func (f scriptFilterNodeFactory) Labels() []string { return []string{"True", "False", "Failure"} }
func (f scriptFilterNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &scriptFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *scriptFilterNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")
	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")
	failureLabelNode := n.GetLinkedNode("Failure")

	scriptEngine := NewScriptEngine(*msg, "Filter", n.Script)
	isTrue, err := scriptEngine.ScriptOnFilter()
	if err != nil {
		if failureLabelNode != nil {
			n.Debug(msg, message.DEBUGOUT, err.Error())
			return failureLabelNode.Handle(msg)
		}
	}

	if isTrue == true && trueLabelNode != nil {
		n.Debug(msg, message.DEBUGOUT, "")
		return trueLabelNode.Handle(msg)
	} else {
		if falseLabelNode != nil {
			n.Debug(msg, message.DEBUGOUT, "Script脚本执行失败")
			return falseLabelNode.Handle(msg)
		}
	}

	return nil
}
