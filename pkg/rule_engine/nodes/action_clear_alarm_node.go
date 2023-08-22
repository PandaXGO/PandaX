package nodes

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"pandax/apps/device/services"
	"pandax/pkg/global"
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

	alarm := services.DeviceAlarmModelDao.FindOneByType(msg.GetMetadata().GetKeyValue("deviceId").(string), n.AlarmType, "0")
	if alarm.DeviceId != "" {
		log.Println("清除告警")
		alarm.State = global.CLEARED
		marshal, _ := json.Marshal(msg.GetMsg())
		alarm.Details = string(marshal)
		err := services.DeviceAlarmModelDao.Update(*alarm)
		if err != nil {
			if failure != nil {
				return failure.Handle(msg)
			}
		} else {
			if cleared != nil {
				return cleared.Handle(msg)
			}
		}
	} else {
		if failure != nil {
			return failure.Handle(msg)
		}
	}
	return nil
}
