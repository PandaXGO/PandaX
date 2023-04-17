package nodes

import (
	"github.com/sirupsen/logrus"
	"log"
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
	return []string{
		"Failure", "True", "False",
		message.EventAttributesType,
		message.EventAlarmType,
		message.EventTelemetryType,
		message.EventUpEventType,
		message.EventConnectType,
		message.EventDisConnectType,
	}
}
func (f switchFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &switchFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *switchFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	scriptEngine := NewScriptEngine()
	SwitchResults, err := scriptEngine.ScriptOnSwitch(msg, n.Script)
	log.Println("开始执行switchFilterNode", SwitchResults)
	if err != nil {
		return nil
	}
	nodes := n.GetLinkedNodes()
	for label, node := range nodes {
		for _, switchresult := range SwitchResults {
			if label == switchresult {
				node.Handle(msg)
			}
		}
	}
	return nil
}
