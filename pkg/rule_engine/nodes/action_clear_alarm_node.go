package nodes

import (
	"encoding/json"
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
func (f clearAlarmNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &clearAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *clearAlarmNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	cleared := n.GetLinkedNode("Cleared")
	failure := n.GetLinkedNode("Failure")
	var err error
	alarm, err := services.DeviceAlarmModelDao.FindOneByType(msg.Metadata.GetValue("deviceId").(string), n.AlarmType, "0")
	if err == nil {
		alarm.State = global.CLEARED
		marshal, _ := json.Marshal(msg.Msg)
		alarm.Details = string(marshal)
		err = services.DeviceAlarmModelDao.Update(*alarm)
		if err == nil {
			if cleared != nil {
				n.Debug(msg, message.DEBUGOUT, "")
				return cleared.Handle(msg)
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
