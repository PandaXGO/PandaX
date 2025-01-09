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
    msgData := make(map[string]any)
    // 去掉多余的参数
    if pid, ok := msg.Metadata.GetValue("productId").(string); ok {
        pts, err := services.ProductTemplateModelDao.FindList(entity.ProductTemplate{Pid: pid, Classify: strings.ToLower(message.TelemetrgMes)})
        if err != nil {
            return errors.New("为获取到设备物模型信息")
        }
        for _, pt := range *pts {
            value := msg.Msg.GetValue(pt.key)
            msgData[pt.key] = value
        }
    } else {
        return errors.New("元素组中为获取到设备ID")
    }
    ts := msg.Msg.GetValue("ts")
    if ts == nil {
        msgData["ts"] = time.Now().Local().Format("2006-01-02 15:04:05.000")
    } else {
        msgData["ts"] = ts
    }
	err := global.TdDb.InsertDevice(deviceName+"_telemetry", msgData)
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
