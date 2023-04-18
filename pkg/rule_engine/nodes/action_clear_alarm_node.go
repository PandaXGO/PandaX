package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

const ClearAlarmNodeName = "ClearAlarmNode"

type clearAlarmNodeFactory struct{}

type clearAlarmNode struct {
	bareNode
	Script    string `json:"script" yaml:"script"`
	AlarmType string `json:"alarmType" yaml:"alarmType"`
}

func (f clearAlarmNodeFactory) Name() string     { return ClearAlarmNodeName }
func (f clearAlarmNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f clearAlarmNodeFactory) Labels() []string { return []string{"Cleared", "Failure"} }
func (f clearAlarmNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &clearAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *clearAlarmNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	cleared := n.GetLinkedNode("Cleared")
	failure := n.GetLinkedNode("Failure")

	scriptEngine := NewScriptEngine(msg, "Details", n.Script)
	details, err := scriptEngine.ScriptAlarmDetails()
	if err != nil {
		return failure.Handle(msg)
	}
	// TODO 编写创建告警信息
	logrus.Info(details)
	cleared.Handle(msg)
	return nil
}
