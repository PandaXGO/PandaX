package util

import (
	"context"
	"encoding/json"
	"errors"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	ruleService "pandax/apps/rule/services"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/tool"
)

func BuildRunDeviceRpc(deviceId, mode string, metadata map[string]interface{}) error {
	one := services.DeviceModelDao.FindOne(deviceId)
	if one.LinkStatus != global.ONLINE {
		return errors.New("设备不在线无法设置属性")
	}
	findOne := ruleService.RuleChainModelDao.FindOne(one.Product.RuleChainId)
	ruleData := ruleEntity.RuleDataJson{}
	err := tool.StringToStruct(findOne.RuleDataJson, &ruleData)
	if err != nil {
		global.Log.Error("规则链数据转化失败", err)
		return errors.New("规则链数据转化失败")
	}
	dataCode := ruleData.LfData.DataCode
	code, _ := json.Marshal(dataCode)
	//新建规则链实体
	instance, errs := rule_engine.NewRuleChainInstance(findOne.Id, code)
	if len(errs) > 0 {
		return errs[0]
	}
	metadataVals := map[string]interface{}{
		"deviceId":       one.Id,
		"mode":           mode,
		"deviceName":     one.Name,
		"deviceType":     one.DeviceType,
		"deviceProtocol": one.Product.ProtocolName,
		"productId":      one.Pid,
		"orgId":          one.OrgId,
		"owner":          one.Owner,
	}
	msg := message.NewMessage(one.Owner, message.RpcRequestToDevice, metadata, metadataVals)
	err = instance.StartRuleChain(context.Background(), msg)
	if err != nil {
		global.Log.Error("规则链执行失败", errs)
	}
	return err
}
