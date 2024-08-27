package nodes

import (
	"errors"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
)

type saveTimeSeriesNode struct {
	bareNode
}

type saveTimeSeriesNodeFactory struct{}

func (f saveTimeSeriesNodeFactory) Name() string     { return "SaveTimeSeriesNode" }
func (f saveTimeSeriesNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f saveTimeSeriesNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f saveTimeSeriesNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &saveTimeSeriesNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *saveTimeSeriesNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	/*	if msg.MsgType != message.TelemetryMes && msg.MsgType != message.RowMes{
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return nil
		}
	}*/
	var deviceName string
	if dn, ok := msg.Metadata.GetValue("deviceName").(string); ok {
		deviceName = dn
	} else {
		return errors.New("元数据中为获取到设备ID")
	}
	err := global.TdDb.InsertDevice(deviceName+"_telemetry", msg.Msg)
	if err != nil {
		n.Debug(msg, message.DEBUGOUT, err.Error())
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLabelNode != nil {
		n.Debug(msg, message.DEBUGOUT, "")
		return successLabelNode.Handle(msg)
	}
	return nil
}
