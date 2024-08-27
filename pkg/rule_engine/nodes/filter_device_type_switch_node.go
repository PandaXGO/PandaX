package nodes

import (
	"pandax/pkg/rule_engine/message"
)

// 检查关联关系
// 该消息来自与哪个实体或到那个实体
type deviceTypeSwitchNode struct {
	bareNode
}

type deviceTypeSwitchNodeFactory struct{}

func (f deviceTypeSwitchNodeFactory) Name() string     { return "DeviceTypeSwitchNode" }
func (f deviceTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f deviceTypeSwitchNodeFactory) Labels() []string {
	return []string{message.DEVICE, message.GATEWAY}
}
func (f deviceTypeSwitchNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &deviceTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *deviceTypeSwitchNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	deviceLabelNode := n.GetLinkedNode(message.DEVICE)
	gatewayLabelNode := n.GetLinkedNode(message.GATEWAY)
	deviceType := msg.Metadata.GetStringValue("deviceType")
	if deviceType == message.DEVICE {
		if deviceLabelNode != nil {
			n.Debug(msg, message.DEBUGOUT, "")
			return deviceLabelNode.Handle(msg)
		}
	} else if deviceType == message.GATEWAY {
		if gatewayLabelNode != nil {
			n.Debug(msg, message.DEBUGOUT, "")
			return gatewayLabelNode.Handle(msg)
		}
	}
	n.Debug(msg, message.DEBUGOUT, "没有匹配的设备类型")
	return nil
}
