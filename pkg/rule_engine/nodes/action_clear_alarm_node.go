package nodes

import (
	"pandax/pkg/rule_engine/message"
	"time"

	"github.com/sirupsen/logrus"
)

const ClearAlarmNodeName = "ClearAlarmNode"

type clearAlarmNodeFactory struct{}

type clearAlarmNode struct {
	bareNode
	DetailBuilderScript string     `json:"detailBuilderScript" yaml:"detailBuilderScript"`
	AlarmType           string     `json:"alarmType" yaml:"alarmType"`
	AlarmSeverity       string     `json:"alarmSeverity" yaml:"alarmSeverity"`
	Propagate           string     `json:"propagate" yaml:"propagate"`
	AlarmStartTime      *time.Time `json:"alarmStartTime" yaml:"alarmStartTime"`
	AlarmEndTime        *time.Time `json:"alarmEndTime" yaml:"alarmEndTime"`
}

func (f clearAlarmNodeFactory) Name() string     { return ClearAlarmNodeName }
func (f clearAlarmNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f clearAlarmNodeFactory) Labels() []string { return []string{"Created", "Updated", "Failure"} }
func (f clearAlarmNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &clearAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *clearAlarmNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	return nil
}
