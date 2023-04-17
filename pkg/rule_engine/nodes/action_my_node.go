package nodes

import (
	"log"
	"pandax/pkg/rule_engine/message"
)

type MyNode struct {
	bareNode
}

type myNodeFactory struct{}

func (f myNodeFactory) Name() string     { return "MyNode" }
func (f myNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f myNodeFactory) Labels() []string { return []string{"Next", "Next2"} }
func (f myNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &MyNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *MyNode) Handle(msg message.Message) error {
	nextLableNode := n.GetLinkedNode("Next")

	log.Println(nextLableNode.Name())
	return nil
}
