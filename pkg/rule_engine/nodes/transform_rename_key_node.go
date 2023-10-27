package nodes

import (
	"pandax/pkg/rule_engine/message"
)

type transformRenameKeyNode struct {
	bareNode
	FormType string    `json:"formType" yaml:"formType"` //msg metadata
	Keys     []KeyName `json:"keys" yaml:"keys"`
}
type KeyName struct {
	OldName string `json:"oldName" yaml:"oldName"`
	NewName string `json:"newName" yaml:"newName"`
}
type transformRenameKeyNodeFactory struct{}

func (f transformRenameKeyNodeFactory) Name() string     { return "RenameKeyNode" }
func (f transformRenameKeyNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }
func (f transformRenameKeyNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f transformRenameKeyNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &transformRenameKeyNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *transformRenameKeyNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if n.FormType == "msg" {
		data := msg.Msg
		for _, key := range n.Keys {
			if _, found := data[key.OldName]; found {
				data[key.NewName] = data[key.OldName]
				delete(data, key.OldName)
				msg.Msg = data
			}
		}
	} else if n.FormType == "metadata" {
		data := msg.Metadata
		for _, key := range n.Keys {
			if data.GetValue(key.OldName) != nil {
				data[key.NewName] = data[key.OldName]
				delete(data, key.OldName)
				msg.Metadata = data
			}
		}
	} else {
		if failureLabelNode != nil {
			n.Debug(msg, message.DEBUGOUT, "未识别FormType")
			return failureLabelNode.Handle(msg)
		}
	}
	if successLabelNode != nil {
		n.Debug(msg, message.DEBUGOUT, "")
		return successLabelNode.Handle(msg)
	}
	return nil
}
