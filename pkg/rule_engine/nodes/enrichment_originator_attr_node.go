package nodes

type enrichmentOriginatorAttrNodeFactory struct{}

func (f enrichmentOriginatorAttrNodeFactory) Name() string     { return "EnrichmentOriginatorAttribute" }
func (f enrichmentOriginatorAttrNodeFactory) Category() string { return NODE_CATEGORY_ENRICHMENT }
func (f enrichmentOriginatorAttrNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
