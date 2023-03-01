package nodes

type rPCCallRequestNode struct {
	bareNode
	TimeoutInSeconds int
}

type rpcCallRequestNodeFactory struct{}

func (f rpcCallRequestNodeFactory) Name() string     { return "RPCCallRequestNode" }
func (f rpcCallRequestNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f rpcCallRequestNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
