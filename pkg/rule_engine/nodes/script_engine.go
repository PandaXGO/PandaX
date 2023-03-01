package nodes

import "pandax/pkg/rule_engine/message"

type ScriptEngine interface {
	ScriptOnMessage(msg message.Message, script string) (message.Message, error)
	//used by filter_switch_node
	ScriptOnSwitch(msg message.Message, script string) ([]string, error)
	//used by filter_script_node
	ScriptOnFilter(msg message.Message, script string) (bool, error)
	ScriptToString(msg message.Message, script string) (string, error)
}

type baseScriptEngine struct {
}

func NewScriptEngine() ScriptEngine {
	return &baseScriptEngine{}
}

func (bse *baseScriptEngine) ScriptOnMessage(msg message.Message, script string) (message.Message, error) {

	return nil, nil
}

func (bse *baseScriptEngine) ScriptOnSwitch(msg message.Message, script string) ([]string, error) {

	return nil, nil
}

func (bse *baseScriptEngine) ScriptOnFilter(msg message.Message, script string) (bool, error) {

	return false, nil
}

func (bse *baseScriptEngine) ScriptToString(msg message.Message, script string) (string, error) {

	return "", nil
}
