package nodes

// init register all node's factory
func init() {
	RegisterFactory(inputNodeFactory{})
	RegisterFactory(switchFilterNodeFactory{})
	RegisterFactory(scriptFilterNodeFactory{})
	RegisterFactory(messageTypeFilterNodeFactory{})
	RegisterFactory(messageTypeSwitchNodeFactory{})
	RegisterFactory(deviceTypeSwitchNodeFactory{})

	RegisterFactory(transformDeleteKeyNodeFactory{})
	RegisterFactory(transformRenameKeyNodeFactory{})
	RegisterFactory(transformScriptNodeFactory{})

	RegisterFactory(createAlarmNodeFactory{})
	RegisterFactory(clearAlarmNodeFactory{})
	RegisterFactory(messageGeneratorNodeFactory{})
	RegisterFactory(logNodeFactory{})
	RegisterFactory(saveAttributesNodeFactory{})
	RegisterFactory(saveTimeSeriesNodeFactory{})
	RegisterFactory(delayNodeFactory{})

	RegisterFactory(externalDingNodeFactory{})
	RegisterFactory(externalWechatNodeFactory{})
	RegisterFactory(externalKafkaNodeFactory{})
	RegisterFactory(externalNatsNodeFactory{})
	RegisterFactory(externalMqttNodeFactory{})
	RegisterFactory(externalRestapiNodeFactory{})
	RegisterFactory(externalSendEmailNodeFactory{})
	RegisterFactory(externalSendSmsNodeFactory{})
	RegisterFactory(externalRuleChainNodeFactory{})

}
