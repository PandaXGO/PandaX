package nodes

// init register all node's factory
func init() {
	RegisterFactory(inputNodeFactory{})
	RegisterFactory(switchFilterNodeFactory{})
	RegisterFactory(delayNodeFactory{})
}
