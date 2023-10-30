package nodes

import (
	"pandax/apps/device/services"
	"pandax/pkg/rule_engine/message"
)

type transformMetadataNode struct {
	bareNode
	DeviceId string `json:"deviceId" yaml:"deviceId"`
}

type transformMetadataNodeFactory struct{}

func (f transformMetadataNodeFactory) Name() string     { return "MetadataKeyNode" }
func (f transformMetadataNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }
func (f transformMetadataNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f transformMetadataNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &transformMetadataNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *transformMetadataNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	// 获取设备信息
	device, err := services.DeviceModelDao.FindOne(n.DeviceId)
	if err != nil {
		n.Debug(msg, message.DEBUGOUT, err.Error())
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return err
		}
	}
	// 更改元数据基本信息
	msg.Metadata = map[string]interface{}{
		"deviceId":       n.DeviceId,
		"deviceName":     device.Name,
		"deviceType":     device.DeviceType,
		"deviceProtocol": device.Product.ProtocolName,
		"productId":      device.Pid,
		"orgId":          device.OrgId,
		"owner":          device.Owner,
	}
	if successLabelNode != nil {
		n.Debug(msg, message.DEBUGOUT, "")
		return successLabelNode.Handle(msg)
	}
	return nil
}
