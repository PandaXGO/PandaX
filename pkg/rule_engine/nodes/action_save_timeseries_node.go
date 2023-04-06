package nodes

import (
	"pandax/pkg/rule_engine/message"
)

type SaveTimeSeriesNode struct {
	bareNode
}

type saveTimeSeriesNodeFactory struct{}

func (f saveTimeSeriesNodeFactory) Name() string     { return "SaveTimeSeriesNode" }
func (f saveTimeSeriesNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f saveTimeSeriesNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f saveTimeSeriesNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}

func (n *SaveTimeSeriesNode) Handle(msg message.Message) error {
	successLableNode := n.GetLinkedNode("Success")
	//failureLableNode := n.GetLinkedNode("Failure")

	return successLableNode.Handle(msg)
}
