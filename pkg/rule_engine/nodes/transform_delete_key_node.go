package nodes

import (
	"pandax/pkg/rule_engine/message"
	"strings"
)

type transformDeleteKeyNode struct {
	bareNode
	FormType string `json:"formType" yaml:"formType"` //msg metadata
	Keys     string `json:"keys" yaml:"keys"`
}
type transformDeleteKeyNodeFactory struct{}

func (f transformDeleteKeyNodeFactory) Name() string     { return "DeleteKeyNode" }
func (f transformDeleteKeyNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }
func (f transformDeleteKeyNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f transformDeleteKeyNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &transformDeleteKeyNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *transformDeleteKeyNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	keys := strings.Split(n.Keys, ",")
	if n.FormType == "msg" {
		data := msg.Msg
		for _, key := range keys {
			if _, found := data[key]; found {
				delete(data, key)
				msg.Msg = data
			}
		}
	} else if n.FormType == "metadata" {
		data := msg.Metadata
		for _, key := range keys {
			if data.GetValue(key) != nil {
				delete(data, key)
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
