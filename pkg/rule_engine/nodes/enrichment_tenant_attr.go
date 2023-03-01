package nodes

type enrichmentTenantNode struct {
	bareNode
}

type enrichmentTenantNodeFactory struct{}

func (f enrichmentTenantNodeFactory) Name() string     { return "EnrichmentTenantNode" }
func (f enrichmentTenantNodeFactory) Category() string { return NODE_CATEGORY_ENRICHMENT }
func (f enrichmentTenantNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}
	node := &enrichmentTenantNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}
