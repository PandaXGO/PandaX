package nodes

import (
	"dz-iot-server/rule_engine/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	DEVICE  = "DEVICE"
	GATEWAY = "GATEWAY"
)

//检查关联关系
//该消息来自与哪个实体或到那个实体
type deviceTypeSwitchNode struct {
	bareNode
}

type deviceTypeSwitchNodeFactory struct{}

func (f deviceTypeSwitchNodeFactory) Name() string     { return "DeviceTypeSwitch" }
func (f deviceTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f deviceTypeSwitchNodeFactory) Labels() []string { return []string{DEVICE, GATEWAY} }
func (f deviceTypeSwitchNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &deviceTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *deviceTypeSwitchNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	deviceLabelNode := n.GetLinkedNode(DEVICE)
	gatewayLabelNode := n.GetLinkedNode(GATEWAY)

	if deviceLabelNode == nil && gatewayLabelNode == nil {
		return fmt.Errorf("no device and gateway label linked node in %s", n.Name())
	}

	if msg.GetMetadata().GetKeyValue("deviceType") == DEVICE {
		return deviceLabelNode.Handle(msg)
	}
	return gatewayLabelNode.Handle(msg)
}
