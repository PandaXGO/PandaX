package nodes

import (
	"dz-iot-server/rule_engine/message"
	"fmt"

	"github.com/sirupsen/logrus"
)

type createAlarmNode struct {
	bareNode
	DetailBuilderScript string `json:"detailBuilderScript" yaml:"detailBuilderScript"`
	AlarmType           string `json:"alarmType" yaml:"alarmType"`
	AlarmSeverity       string `json:"alarmSeverity" yaml:"alarmSeverity"`
	Propagate           string `json:"propagate" yaml:"propagate"`
	AlarmStartTime      string `json:"alarmStartTime" yaml:"alarmStartTime"`
	AlarmEndTime        string `json:"alarmEndTime" yaml:"alarmEndTime"`
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
	node2 := n.GetLinkedNode("Updated")
	node3 := n.GetLinkedNode("Failure")
	if node1 == nil || node2 == nil || node3 == nil {
		return fmt.Errorf("no valid label linked node in %s", n.Name())
	}

	return nil
}
