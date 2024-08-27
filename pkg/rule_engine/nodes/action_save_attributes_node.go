package nodes

import (
	"errors"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
)

type saveAttributesNode struct {
	bareNode
}

type saveAttributesNodeFactory struct{}

func (f saveAttributesNodeFactory) Name() string     { return "SaveAttributesNode" }
func (f saveAttributesNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f saveAttributesNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f saveAttributesNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &saveAttributesNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *saveAttributesNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")
	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	/*if msg.MsgType != message.AttributesMes {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return nil
		}
	}*/
	//deviceId := msg.GetMetadata().GetValues()["deviceId"].(string)
	var deviceName string
	if dn, ok := msg.Metadata.GetValue("deviceName").(string); ok {
		deviceName = dn
	} else {
		return errors.New("元数据中为获取到设备ID")
	}
	err := global.TdDb.InsertDevice(deviceName+"_attributes", msg.Msg)
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
