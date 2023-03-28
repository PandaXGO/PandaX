package rule_engine

import (
	"io/ioutil"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"
	"testing"
)

func TestNewRuleChainInstance(t *testing.T) {
	buf, err := ioutil.ReadFile("./manifest/manifest_sample.json")
	if err != nil {
		t.Error(err)
	}
	_, errs := NewRuleChainInstance(buf)
	if len(errs) > 0 {
		t.Error(errs[0])
	}
}

func TestScriptEngine(t *testing.T) {
	metadata := message.NewDefaultMetadata(map[string]interface{}{"device": "aa"})
	msg := message.NewMessageWithDetail("1", message.MessageTypeConnectEvent, map[string]interface{}{"aa": 5}, metadata)
	scriptEngine := nodes.NewScriptEngine()
	const script = `
    function Switch(msg, metadata, msgType) {
        function nextRelation(metadata, msg) {
			return ['one','nine'];
		}
		if(metadata.device === 'aa') {
			return ['six'];
		}
		if(msgType === 'Post telemetry') {
			return ['two'];
		}
		return nextRelation(metadata, msg);
	}
    `
	SwitchResults, err := scriptEngine.ScriptOnSwitch(msg, script)

	if err != nil {
		t.Error(err)
	}
	t.Log(SwitchResults)
}

func TestScriptOnMessage(t *testing.T) {
	metadata := message.NewDefaultMetadata(map[string]interface{}{"device": "aa"})
	msg := message.NewMessageWithDetail("1", message.MessageTypeConnectEvent, map[string]interface{}{"aa": 5}, metadata)
	scriptEngine := nodes.NewScriptEngine()
	const script = `
    function Transform(msg, metadata, msgType) {
        msg.bb = "33"
		metadata.event = 55
		return {msg: msg, metadata: metadata, msgType: msgType};
	}
    `
	ScriptOnMessageResults, err := scriptEngine.ScriptOnMessage(msg, script)

	if err != nil {
		t.Error(err)
	}
	t.Log(ScriptOnMessageResults.GetMetadata())
}
