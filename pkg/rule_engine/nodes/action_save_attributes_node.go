package nodes

import (
	"dz-iot-server/rule_engine/message"
	"fmt"
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
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")
	if successLableNode == nil || failureLableNode == nil {
		return fmt.Errorf("no valid label linked node in %s", n.Name())
	}
	if msg.GetType() != "POST_ATTRIBUTES_REQUEST" {
		return failureLableNode.Handle(msg)
	}

	return nil
}
