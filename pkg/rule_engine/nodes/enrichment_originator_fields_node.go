package nodes

type enrichmentOriginatorFieldsNodeFactory struct{}

func (f enrichmentOriginatorFieldsNodeFactory) Name() string     { return "EnrichmentOriginatorFieldsNode" }
func (f enrichmentOriginatorFieldsNodeFactory) Category() string { return NODE_CATEGORY_ENRICHMENT }
func (f enrichmentOriginatorFieldsNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
