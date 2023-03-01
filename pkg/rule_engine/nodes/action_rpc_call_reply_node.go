package nodes

type rpcCallReplyNodeFactory struct{}

func (f rpcCallReplyNodeFactory) Name() string     { return "RPCCallReplyNode" }
func (f rpcCallReplyNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f rpcCallReplyNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
