package nodes

type deleteRelationNodeFactory struct{}

func (f deleteRelationNodeFactory) Name() string     { return "DeleteRelationNode" }
func (f deleteRelationNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f deleteRelationNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
