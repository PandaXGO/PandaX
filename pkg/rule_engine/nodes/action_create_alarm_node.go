package nodes

import (
	"encoding/json"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/pkg/global"
	"pandax/pkg/global_model"
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
func (f createAlarmNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &createAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *createAlarmNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	created := n.GetLinkedNode("Created")
	updated := n.GetLinkedNode("Updated")
	failure := n.GetLinkedNode("Failure")
	var err error
	alarm, err := services.DeviceAlarmModelDao.FindOneByType(msg.Metadata.GetValue("deviceId").(string), n.AlarmType, "0")
	if err == nil {
		marshal, _ := json.Marshal(msg.Msg)
		alarm.Details = string(marshal)
		err = services.DeviceAlarmModelDao.Update(*alarm)
		if err == nil {
			if updated != nil {
				n.Debug(msg, message.DEBUGOUT, "")
				return updated.Handle(msg)
			}
		}
	} else {
		alarm = &entity.DeviceAlarm{}
		alarm.Id = global_model.GenerateID()
		alarm.DeviceId = msg.Metadata.GetValue("deviceId").(string)
		alarm.ProductId = msg.Metadata.GetValue("productId").(string)
		alarm.Name = msg.Metadata.GetValue("deviceName").(string)
		alarm.Level = n.AlarmSeverity
		alarm.State = global.ALARMING
		alarm.Type = n.AlarmType
		alarm.Time = time.Now()
		alarm.OrgId = msg.Metadata.GetValue("orgId").(int64)
		alarm.Owner = msg.Metadata.GetValue("owner").(string)
		marshal, _ := json.Marshal(msg.Msg)
		alarm.Details = string(marshal)
		err = services.DeviceAlarmModelDao.Insert(*alarm)
		if err == nil {
			if created != nil {
				n.Debug(msg, message.DEBUGOUT, "")
				return created.Handle(msg)
			}
		}
	}
	if err != nil {
		n.Debug(msg, message.DEBUGOUT, err.Error())
		if failure != nil {
			return failure.Handle(msg)
		} else {
			return err
		}
	}
	return nil
}
