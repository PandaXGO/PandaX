package nodes

type assignToCustomerNode struct {
	bareNode
}

type assignToCustomerFactory struct{}

func (f assignToCustomerFactory) Name() string     { return "AssignCustomerFactoryNode" }
func (f assignToCustomerFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f assignToCustomerFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
