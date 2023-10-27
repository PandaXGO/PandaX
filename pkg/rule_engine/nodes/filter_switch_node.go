package nodes

import (
	"pandax/pkg/rule_engine/message"
)

type switchFilterNode struct {
	bareNode
	Script string `json:"script" yaml:"script"`
}

type switchFilterNodeFactory struct{}

func (f switchFilterNodeFactory) Name() string     { return "SwitchNode" }
func (f switchFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f switchFilterNodeFactory) Labels() []string {
	return []string{"Failure"}
}
func (f switchFilterNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &switchFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *switchFilterNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	failureLabelNode := n.GetLinkedNode("Failure")

	scriptEngine := NewScriptEngine(*msg, "Switch", n.Script)
	SwitchResults, err := scriptEngine.ScriptOnSwitch()
	if err != nil {
		n.Debug(msg, message.DEBUGOUT, err.Error())
		return failureLabelNode.Handle(msg)
	}
	nodes := n.GetLinkedNodes()
	for label, node := range nodes {
		for _, switchresult := range SwitchResults {
			if label == switchresult {
				go node.Handle(msg)
			}
		}
	}
	n.Debug(msg, message.DEBUGOUT, "")
	return nil
}
