package rule_engine

import (
	"context"
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
	instance, errs := NewRuleChainInstance(buf)
	if len(errs) > 0 {
		t.Error(errs[0])
	}

	metadata := message.NewDefaultMetadata(map[string]interface{}{
		"deviceName": "ws432",
		"deviceId":   "d_1928b99619910dae5a001fa7",
		"deviceType": "direct",
		"productId":  "p_3ba460634520cf4590dc90e5",
	})
	msg := message.NewMessageWithDetail("1", message.TelemetryMes, map[string]interface{}{"temperature": 60.4, "humidity": 32.5}, metadata)
	t.Log("开始执行力流程")
	err = instance.StartRuleChain(context.Background(), msg)
	if err != nil {
		t.Log(err)
	}
}

func TestScriptEngine(t *testing.T) {
	metadata := message.NewDefaultMetadata(map[string]interface{}{"device": "aa"})
	msg := message.NewMessageWithDetail("1", message.UpEventMes, map[string]interface{}{"aa": 5}, metadata)
	const baseScript = `
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
    `
	scriptEngine := nodes.NewScriptEngine(msg, "Switch", baseScript)
	SwitchResults, err := scriptEngine.ScriptOnSwitch()

	if err != nil {
		t.Error(err)
	}
	t.Log(SwitchResults)
}

func TestScriptOnMessage(t *testing.T) {
	metadata := message.NewDefaultMetadata(map[string]interface{}{"device": "aa"})
	msg := message.NewMessageWithDetail("1", message.UpEventMes, map[string]interface{}{"aa": 5}, metadata)

	const baseScript = `
        msg.bb = "33"
		metadata.event = 55
		return {msg: msg, metadata: metadata, msgType: msgType};
    `
	scriptEngine := nodes.NewScriptEngine(msg, "Transform", baseScript)
	ScriptOnMessageResults, err := scriptEngine.ScriptOnMessage()

	if err != nil {
		t.Error(err)
	}
	t.Log(ScriptOnMessageResults.GetMetadata())
}
