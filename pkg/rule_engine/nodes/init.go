package nodes

// init register all node's factory
func init() {
	RegisterFactory(inputNodeFactory{})
}
