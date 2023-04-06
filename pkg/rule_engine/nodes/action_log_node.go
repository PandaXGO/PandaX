package nodes

import (
	"fmt"
	"log"
	"pandax/pkg/rule_engine/message"
)

type logNode struct {
	bareNode
	Script string `json:"script"`
}

type logNodeFactory struct{}

func (f logNodeFactory) Name() string     { return "LogNode" }
func (f logNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f logNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f logNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &logNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *logNode) Handle(msg message.Message) error {
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")

	scriptEngine := NewScriptEngine()
	logMessage, err := scriptEngine.ScriptToString(msg, n.Script)

	if successLableNode == nil || failureLableNode == nil {
		return fmt.Errorf("no valid label linked node in %s", n.Name())
	}
	if err != nil {
		return failureLableNode.Handle(msg)
	}
	log.Println(logMessage)
	return successLableNode.Handle(msg)
}
