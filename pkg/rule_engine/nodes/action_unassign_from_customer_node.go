package nodes

type unassignFromCustomerNodeFactory struct{}

func (f unassignFromCustomerNodeFactory) Name() string     { return "UnassignFromCustomerNode" }
func (f unassignFromCustomerNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f unassignFromCustomerNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
