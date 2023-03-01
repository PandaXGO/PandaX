package nodes

type enrichmentDeviceAttrNodeFactory struct{}

func (f enrichmentDeviceAttrNodeFactory) Name() string     { return "EnrichmentDeviceAttrbute" }
func (f enrichmentDeviceAttrNodeFactory) Category() string { return NODE_CATEGORY_ENRICHMENT }
func (f enrichmentDeviceAttrNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
