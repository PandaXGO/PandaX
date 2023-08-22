package nodes

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type ScriptEngine interface {
	ScriptOnMessage() (message.Message, error)
	ScriptOnSwitch() ([]string, error)
	ScriptOnFilter() (bool, error)
	ScriptToString() (string, error)
	ScriptGenerate() (map[string]interface{}, error)
}

type baseScriptEngine struct {
	Fun    string
	Script string
	Msg    message.Message
}

func NewScriptEngine(msg message.Message, fun string, script string) ScriptEngine {
	return &baseScriptEngine{
		Fun:    fun,
		Script: fmt.Sprintf("function %s(msg, metadata, msgType) { %s }", fun, script),
		Msg:    msg,
	}
}

func (bse *baseScriptEngine) ScriptOnMessage() (message.Message, error) {
	msg := bse.Msg
	vm := goja.New()
	_, err := vm.RunString(bse.Script)
	if err != nil {
		logrus.Info("JS代码有问题")
		return nil, err
	}
	var fn func(map[string]interface{}, map[string]interface{}, string) map[string]interface{}
	err = vm.ExportTo(vm.Get(bse.Fun), &fn)
	if err != nil {
		logrus.Info("Js函数映射到 Go 函数失败！")
		return nil, err
	}
	datas := fn(msg.GetMsg(), msg.GetMetadata().GetValues(), msg.GetType())
	msg.SetMsg(datas["msg"].(map[string]interface{}))
	msg.SetMetadata(message.NewDefaultMetadata(datas["metadata"].(map[string]interface{})))
	msg.SetType(datas["msgType"].(string))
	return msg, nil
}

func (bse *baseScriptEngine) ScriptOnSwitch() ([]string, error) {
	msg := bse.Msg
	vm := goja.New()
	_, err := vm.RunString(bse.Script)
	if err != nil {
		logrus.Info("JS代码有问题")
		return nil, err
	}
	var fn func(map[string]interface{}, map[string]interface{}, string) []string
	err = vm.ExportTo(vm.Get(bse.Fun), &fn)
	if err != nil {
		logrus.Info("Js函数映射到 Go 函数失败！")
		return nil, err
	}
	datas := fn(msg.GetMsg(), msg.GetMetadata().GetValues(), msg.GetType())
	return datas, nil
}

func (bse *baseScriptEngine) ScriptOnFilter() (bool, error) {
	msg := bse.Msg
	vm := goja.New()
	_, err := vm.RunString(bse.Script)
	if err != nil {
		logrus.Info("JS代码有问题")
		return false, err
	}
	var fn func(map[string]interface{}, map[string]interface{}, string) bool
	err = vm.ExportTo(vm.Get(bse.Fun), &fn)
	if err != nil {
		logrus.Info("Js函数映射到 Go 函数失败！")
		return false, err
	}
	datas := fn(msg.GetMsg(), msg.GetMetadata().GetValues(), msg.GetType())
	return datas, nil
}

func (bse *baseScriptEngine) ScriptToString() (string, error) {
	msg := bse.Msg
	vm := goja.New()
	_, err := vm.RunString(bse.Script)
	if err != nil {
		logrus.Info("JS代码有问题")
		return "", err
	}
	var fn func(map[string]interface{}, map[string]interface{}, string) string
	err = vm.ExportTo(vm.Get(bse.Fun), &fn)
	if err != nil {
		logrus.Info("Js函数映射到 Go 函数失败！")
		return "", err
	}
	data := fn(msg.GetMsg(), msg.GetMetadata().GetValues(), msg.GetType())
	return data, nil
}

func (bse *baseScriptEngine) ScriptGenerate() (map[string]interface{}, error) {
	msg := bse.Msg
	vm := goja.New()
	_, err := vm.RunString(bse.Script)
	if err != nil {
		logrus.Info("JS代码有问题")
		return nil, err
	}
	var fn func(map[string]interface{}, map[string]interface{}, string) map[string]interface{}
	err = vm.ExportTo(vm.Get(bse.Fun), &fn)
	if err != nil {
		logrus.Info("Js函数映射到 Go 函数失败！")
		return nil, err
	}
	datas := fn(msg.GetMsg(), msg.GetMetadata().GetValues(), msg.GetType())
	return datas, nil
}
