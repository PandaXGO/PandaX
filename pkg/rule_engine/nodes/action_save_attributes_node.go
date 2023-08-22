package nodes

import (
	"github.com/sirupsen/logrus"
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
func (f saveAttributesNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &saveAttributesNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *saveAttributesNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if msg.GetType() != message.AttributesMes {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return nil
		}
	}
	//deviceId := msg.GetMetadata().GetValues()["deviceId"].(string)
	deviceName := msg.GetMetadata().GetValues()["deviceName"].(string)
	err := global.TdDb.InsertDevice(deviceName+"_attributes", msg.GetMsg())
	if err != nil {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		}
	}
	if successLabelNode != nil {
		return successLabelNode.Handle(msg)
	}
	return nil
}
