package nodes

import (
	"fmt"
	"log"
	"pandax/pkg/rule_engine/message"
)

type logNode struct {
	bareNode
	Script string
}

type logNodeFactory struct{}

func (f logNodeFactory) Name() string     { return "LogNode" }
func (f logNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f logNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}
	node := &logNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
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
