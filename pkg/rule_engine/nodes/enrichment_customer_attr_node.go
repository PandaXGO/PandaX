package nodes

import "pandax/pkg/rule_engine/message"

type enrichmentCustomerNode struct {
	bareNode
}

type enrichmentCustomerAttrNodeFactory struct{}

func (f enrichmentCustomerAttrNodeFactory) Name() string     { return "EnrichmentCustomerNode" }
func (f enrichmentCustomerAttrNodeFactory) Category() string { return NODE_CATEGORY_ENRICHMENT }
func (f enrichmentCustomerAttrNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}
	node := &enrichmentCustomerNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *enrichmentCustomerNode) Handle(msg message.Message) error {
	return nil
}
