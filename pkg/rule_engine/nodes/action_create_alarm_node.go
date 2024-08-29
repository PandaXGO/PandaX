package nodes

import (
	"encoding/json"
	"errors"
	"github.com/PandaXGO/PandaKit/utils"
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
	var deviceId string
	if did, ok := msg.Metadata.GetValue("deviceId").(string); ok {
		deviceId = did
	} else {
		return errors.New("元数据中为获取到设备ID")
	}
	var err error
	alarm, err := services.DeviceAlarmModelDao.FindOneByType(deviceId, n.AlarmType, "0")
	if err == nil {
		marshal, _ := json.Marshal(msg.Msg)
		alarm.Details = string(marshal)
		err = services.DeviceAlarmModelDao.Update(*alarm)
		if err == nil {
			n.Debug(msg, message.DEBUGOUT, "")
			if updated != nil {
				return updated.Handle(msg)
			}
		}
	} else {
		alarm = &entity.DeviceAlarm{}
		alarm.Id = utils.GenerateID("")
		alarm.DeviceId = msg.Metadata.GetStringValue("deviceId")
		alarm.ProductId = msg.Metadata.GetStringValue("productId")
		alarm.Name = msg.Metadata.GetStringValue("deviceName")
		alarm.Level = n.AlarmSeverity
		alarm.State = global.ALARMING
		alarm.Type = n.AlarmType
		alarm.Time = time.Now()
		alarm.OrgId = int64(msg.Metadata.GetIntValue("orgId"))
		alarm.Owner = msg.Metadata.GetStringValue("owner")
		marshal, _ := json.Marshal(msg.Msg)
		alarm.Details = string(marshal)
		err = services.DeviceAlarmModelDao.Insert(*alarm)
		if err == nil {
			n.Debug(msg, message.DEBUGOUT, "")
			if created != nil {
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
