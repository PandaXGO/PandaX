package util

import (
	"encoding/json"
	"errors"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	ruleService "pandax/apps/rule/services"
	devicerpc "pandax/pkg/device_rpc"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/tool"
)

func BuildRunDeviceRpc(deviceId, mode string, rp devicerpc.RpcPayload) error {
	device, err := services.DeviceModelDao.FindOne(deviceId)
	if err != nil {
		return err
	}
	if device.LinkStatus != global.ONLINE {
		return errors.New("设备不在线无法设置属性")
	}

	//新建规则链实体
	ruleInstance := rule_engine.RuleEngine.GetRuleInstance(device.Product.Id)
	if ruleInstance == nil {
		findOne, err := ruleService.RuleChainModelDao.FindOne(device.Product.RuleChainId)
		if err != nil {
			global.Log.Error("查询规则链数据失败", err)
			return errors.New("查询规则链数据失败")
		}
		ruleData := ruleEntity.RuleDataJson{}
		err = tool.StringToStruct(findOne.RuleDataJson, &ruleData)
		if err != nil {
			global.Log.Error("规则链数据转化失败", err)
			return errors.New("规则链数据转化失败")
		}
		dataCode := ruleData.DataCode
		code, err := json.Marshal(dataCode)
		if err != nil {
			global.Log.Error("规则链数据解析失败", err)
			return errors.New("规则链数据解析失败")
		}
		ruleInstance, err = rule_engine.NewRuleChainInstance(findOne.Id, code)
		if err != nil {
			return err
		}
		rule_engine.RuleEngine.SaveRuleInstance(device.Product.Id, ruleInstance)
	}
	metadataVals := map[string]interface{}{
		"deviceId":       device.Id,
		"mode":           mode,
		"deviceName":     device.Name,
		"deviceType":     device.DeviceType,
		"deviceProtocol": device.Product.ProtocolName,
		"productId":      device.Pid,
		"orgId":          device.OrgId,
		"owner":          device.Owner,
	}
	msg := message.NewMessage(device.Owner, message.RpcRequestToDevice, rp.ToMap(), metadataVals)
	err = rule_engine.RuleEngine.StartRuleInstance(ruleInstance, msg)
	if err != nil {
		global.Log.Error("规则链执行失败", err)
	}
	return err
}
