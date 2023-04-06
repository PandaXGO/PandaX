package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type transformDeleteKeyNode struct {
	bareNode
	FormType string   `json:"formType" yaml:"formType"` //msg metadata
	Keys     []string `json:"keys" yaml:"keys"`
}
type transformDeleteKeyNodeFactory struct{}

func (f transformDeleteKeyNodeFactory) Name() string     { return "TransformDeleteKeyNode" }
func (f transformDeleteKeyNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }
func (f transformDeleteKeyNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f transformDeleteKeyNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &transformDeleteKeyNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *transformDeleteKeyNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if n.FormType == "msg" {
		data := msg.GetMsg()
		for _, key := range n.Keys {
			if _, found := data[key]; found {
				delete(data, key)
				msg.SetMsg(data)
			}
		}
	} else if n.FormType == "metadata" {
		data := msg.GetMetadata()
		for _, key := range n.Keys {
			if data.GetKeyValue(key) != nil {
				values := data.GetValues()
				delete(values, key)
				msg.SetMetadata(message.NewDefaultMetadata(values))
			}
		}
	} else {
		failureLabelNode.Handle(msg)
	}

	return successLabelNode.Handle(msg)
}
