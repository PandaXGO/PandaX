package nodes

import (
	"pandax/pkg/rule_engine/message"
)

type SaveAttributesNode struct {
	bareNode
}

type saveAttributesNodeFactory struct{}

func (f saveAttributesNodeFactory) Name() string     { return "SaveAttributesNode" }
func (f saveAttributesNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f saveAttributesNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f saveAttributesNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &SaveAttributesNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *SaveAttributesNode) Handle(msg message.Message) error {
	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if msg.GetType() != message.EventAttributesType {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return nil
		}
	}
	//deviceName := msg.GetMetadata().GetValues()["deviceName"].(string)
	//namespace := msg.GetMetadata().GetValues()["namespace"].(string)
	//marshal, err := json.Marshal(msg.GetMsg())

	//if err != nil {
	//	if failureLabelNode != nil {
	//		return failureLabelNode.Handle(msg)
	//	} else {
	//		return nil
	//	}
	//}

	// todo 添加设备上报参数

	if successLabelNode != nil {
		return successLabelNode.Handle(msg)
	} else {
		return nil
	}

	return nil
}
