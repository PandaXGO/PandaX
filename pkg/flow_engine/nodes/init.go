package nodes

// init register all node's factory
func init() {
	RegisterFactory(inputNodeFactory{})
	RegisterFactory(transformScriptNodeFactory{})

	RegisterFactory(externalDingNodeFactory{})
	RegisterFactory(externalWechatNodeFactory{})
	RegisterFactory(externalSendEmailNodeFactory{})
	RegisterFactory(externalSendSmsNodeFactory{})

}
