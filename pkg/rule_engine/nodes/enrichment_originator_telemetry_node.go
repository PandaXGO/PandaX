package nodes

type enrichmentOriginatorTelemetryNodeFactory struct{}

func (f enrichmentOriginatorTelemetryNodeFactory) Name() string {
	return "EnrichmentOriginatorTelemetryNode"
}
func (f enrichmentOriginatorTelemetryNodeFactory) Category() string { return NODE_CATEGORY_ENRICHMENT }
func (f enrichmentOriginatorTelemetryNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
