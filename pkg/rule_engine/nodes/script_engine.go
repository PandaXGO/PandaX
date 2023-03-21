package nodes

import (
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

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
	vm := goja.New()
	_, err := vm.RunString(script)
	if err != nil {
		logrus.Info("JS代码有问题")
		return nil, err
	}
	var fn func(message.Message, message.Metadata, string) []string
	err = vm.ExportTo(vm.Get("Switch"), &fn)
	if err != nil {
		logrus.Info("Js函数映射到 Go 函数失败！")
		return nil, err
	}
	datas := fn(msg, msg.GetMetadata(), msg.GetType())
	return datas, nil
}

func (bse *baseScriptEngine) ScriptOnFilter(msg message.Message, script string) (bool, error) {

	return false, nil
}

func (bse *baseScriptEngine) ScriptToString(msg message.Message, script string) (string, error) {

	return "", nil
}
