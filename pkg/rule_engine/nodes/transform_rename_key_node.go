package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type transformRenameKeyNode struct {
	bareNode
	FormType string    `json:"formType" yaml:"formType"` //msg metadata
	Keys     []KeyName `json:"keys" yaml:"keys"`
}
type KeyName struct {
	oldName string `json:"oldName" yaml:"oldName"`
	newName string `json:"newName" yaml:"newName"`
}
type transformRenameKeyNodeFactory struct{}

func (f transformRenameKeyNodeFactory) Name() string     { return "TransformRenameKeyNode" }
func (f transformRenameKeyNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }
func (f transformRenameKeyNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f transformRenameKeyNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &transformScriptNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *transformRenameKeyNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if n.FormType == "msg" {
		data := msg.GetMsg()
		for _, key := range n.Keys {
			if _, found := data[key.oldName]; found {
				data[key.newName] = data[key.oldName]
				delete(data, key.oldName)
				msg.SetMsg(data)
			}
		}
	} else if n.FormType == "metadata" {
		data := msg.GetMetadata()
		for _, key := range n.Keys {
			if data.GetKeyValue(key.oldName) != nil {
				values := data.GetValues()
				values[key.newName] = values[key.oldName]
				delete(values, key.oldName)
				msg.SetMetadata(message.NewDefaultMetadata(values))
			}
		}
	} else {
		failureLabelNode.Handle(msg)
	}

	return successLabelNode.Handle(msg)
}
