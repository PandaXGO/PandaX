package util

import (
	"context"
	"encoding/json"
	"errors"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	ruleService "pandax/apps/rule/services"
	"pandax/pkg/cache"
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
	dataCode := ruleData.LfData.DataCode
	code, _ := json.Marshal(dataCode)
	//新建规则链实体
	instance := &rule_engine.RuleChainInstance{}
	ruleInstance, bo := cache.GetProductRule(device.Product.Id)
	if !bo {
		instance, err = rule_engine.NewRuleChainInstance(findOne.Id, code)
		if err != nil {
			return err
		}
	} else {
		if data, ok := ruleInstance.(*rule_engine.RuleChainInstance); ok {
			instance = data
		} else {
			return errors.New("规则实体解析错误")
		}
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
	err = instance.StartRuleChain(context.Background(), msg)
	if err != nil {
		global.Log.Error("规则链执行失败", err)
	}
	return err
}
