package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type switchFilterNode struct {
	bareNode
	Scripts string `json:"scripts" yaml:"scripts"`
}

type switchFilterNodeFactory struct{}

func (f switchFilterNodeFactory) Name() string     { return "SwitchNode" }
func (f switchFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f switchFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{}
	node := &switchFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *switchFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	/*scriptEngine := NewScriptEngine()
	SwitchResults, err := scriptEngine.ScriptOnSwitch(msg, n.Scripts)
	if err != nil {
		return nil
	}
	nodes := n.GetLinkedNodes()
	for label, node := range nodes {
		for _, switchresult := range SwitchResults {
			if label == switchresult {
				return node.Handle(msg)
			}
		}
	}*/
	return nil
}
