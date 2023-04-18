package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

//检查关联关系
//该消息来自与哪个实体或到那个实体
type deviceTypeSwitchNode struct {
	bareNode
}

type deviceTypeSwitchNodeFactory struct{}

func (f deviceTypeSwitchNodeFactory) Name() string     { return "DeviceTypeSwitchNode" }
func (f deviceTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f deviceTypeSwitchNodeFactory) Labels() []string {
	return []string{message.DEVICE, message.GATEWAY}
}
func (f deviceTypeSwitchNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &deviceTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *deviceTypeSwitchNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	deviceLabelNode := n.GetLinkedNode(message.DEVICE)
	gatewayLabelNode := n.GetLinkedNode(message.GATEWAY)

	if msg.GetMetadata().GetKeyValue("deviceType") == message.DEVICE {
		if deviceLabelNode != nil {
			return deviceLabelNode.Handle(msg)
		}
	}
	if msg.GetMetadata().GetKeyValue("deviceType") == message.GATEWAY {
		if gatewayLabelNode != nil {
			return gatewayLabelNode.Handle(msg)
		}
	}
	return nil
}
