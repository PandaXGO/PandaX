package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type createAlarmNode struct {
	bareNode
	Script        string `json:"script" yaml:"script"`
	AlarmType     string `json:"alarmType" yaml:"alarmType"`
	AlarmSeverity int64  `json:"alarmSeverity" yaml:"alarmSeverity"`
	Propagate     string `json:"propagate" yaml:"propagate"`
	/*	AlarmStartTime      string `json:"alarmStartTime" yaml:"alarmStartTime"`
		AlarmEndTime        string `json:"alarmEndTime" yaml:"alarmEndTime"`*/
}

type createAlarmNodeFactory struct{}

func (f createAlarmNodeFactory) Name() string     { return "CreateAlarmNode" }
func (f createAlarmNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f createAlarmNodeFactory) Labels() []string { return []string{"Created", "Updated", "Failure"} }
func (f createAlarmNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &createAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *createAlarmNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	node1 := n.GetLinkedNode("Created")
	//node2 := n.GetLinkedNode("Updated")
	node3 := n.GetLinkedNode("Failure")

	scriptEngine := NewScriptEngine(msg, "Details", n.Script)
	details, err := scriptEngine.ScriptAlarmDetails()
	if err != nil {
		if node3 != nil {
			return node3.Handle(msg)
		}
	}
	// TODO 创建告警
	logrus.Info(details)
	node1.Handle(msg)

	return nil
}
