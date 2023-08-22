package nodes

import (
	"encoding/json"
	"github.com/kakuilan/kgo"
	"github.com/sirupsen/logrus"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
	"time"
)

type createAlarmNode struct {
	bareNode
	AlarmType     string `json:"alarmType" yaml:"alarmType"`
	AlarmSeverity string `json:"alarmSeverity" yaml:"alarmSeverity"`
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

	created := n.GetLinkedNode("Created")
	updated := n.GetLinkedNode("Updated")
	failure := n.GetLinkedNode("Failure")

	alarm := services.DeviceAlarmModelDao.FindOneByType(msg.GetMetadata().GetKeyValue("deviceId").(string), n.AlarmType, "0")
	if alarm.DeviceId != "" {
		marshal, _ := json.Marshal(msg.GetMsg())
		alarm.Details = string(marshal)
		err := services.DeviceAlarmModelDao.Update(*alarm)
		if err != nil {
			if failure != nil {
				return failure.Handle(msg)
			}
		} else {
			if updated != nil {
				return updated.Handle(msg)
			}
		}
	} else {
		alarm = &entity.DeviceAlarm{}
		alarm.Id = kgo.KStr.Uniqid("a")
		alarm.DeviceId = msg.GetMetadata().GetKeyValue("deviceId").(string)
		alarm.ProductId = msg.GetMetadata().GetKeyValue("productId").(string)
		alarm.Name = msg.GetMetadata().GetKeyValue("deviceName").(string)
		alarm.Level = n.AlarmSeverity
		alarm.State = global.ALARMING
		alarm.Type = n.AlarmType
		alarm.Time = time.Now()
		marshal, _ := json.Marshal(msg.GetMsg())
		alarm.Details = string(marshal)
		err := services.DeviceAlarmModelDao.Insert(*alarm)
		if err != nil {
			if failure != nil {
				return failure.Handle(msg)
			}
		} else {
			if created != nil {
				return created.Handle(msg)
			}
		}
	}
	return nil
}
