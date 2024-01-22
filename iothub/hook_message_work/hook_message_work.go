package hook_message_work

import (
	"context"
	"encoding/json"
	"fmt"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	ruleService "pandax/apps/rule/services"
	"pandax/iothub/netbase"
	"github.com/PandaXGO/PandaKit/biz"
	"pandax/pkg/cache"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"pandax/pkg/rule_engine"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/shadow"
	"pandax/pkg/tool"
	"pandax/pkg/websocket"
)

// 消息处理模块
func (s *HookService) MessageWork() {
	for msg := range s.MessageCh {
		s.handleOne(msg) // 处理消息
	}
}

func (s *HookService) handleOne(msg *netbase.DeviceEventInfo) {
	if s.Ch != nil { // 用于并发限制
		s.Ch <- struct{}{}
	}
	// 用于做优雅关闭， 主要作用是，程序关闭，将队列中的消息处理完成后在关闭，
	s.Wg.Add(1)
	go func() {
		defer func() {
			s.Wg.Done()
			if s.Ch != nil { // 用于并发限制
				<-s.Ch
			}
		}()
		switch msg.Type {
		case message.RowMes, message.AttributesMes, message.TelemetryMes, message.RpcRequestFromDevice:
			msgVals := make(map[string]interface{})
			err := json.Unmarshal([]byte(msg.Datas), &msgVals)
			if err != nil {
				global.Log.Error("数据结构解析错误", err)
				return
			}
			// 发送websocket到云组态
			if msg.Type == message.TelemetryMes {
				go SendZtWebsocket(msg.DeviceId, msg.Datas)
			}
			// 获取规则链代码实体
			instance := getRuleChainInstance(msg.DeviceAuth)
			if instance == nil {
				return
			}
			ruleMessage := buildRuleMessage(msg.DeviceAuth, msgVals, msg.Type)
			err = instance.StartRuleChain(context.Background(), ruleMessage)
			if err != nil {
				global.Log.Error("规则链执行失败", err)
				return
			}
			// 保存设备影子
			if msg.Type != message.RpcRequestFromDevice {
				SetDeviceShadow(msg.DeviceAuth, ruleMessage.Msg, msg.Type)
			}
		case message.DisConnectMes, message.ConnectMes:
			//检测设备影子并修改设备影子状态
			if msg.Type == message.ConnectMes {
				shadow.DeviceShadowInstance.SetOnline(msg.DeviceAuth.Name)
			} else {
				shadow.DeviceShadowInstance.SetOffline(msg.DeviceAuth.Name)
			}
			// 更改设备在线状态
			if msg.Type == message.ConnectMes {
				services.DeviceModelDao.UpdateStatus(msg.DeviceId, global.ONLINE)
			} else {
				services.DeviceModelDao.UpdateStatus(msg.DeviceId, global.OFFLINE)
			}
			// 添加设备连接历史
			data := make(map[string]any)
			err := json.Unmarshal([]byte(msg.Datas), &data)
			if err != nil {
				global.Log.Error("设备连接数据格式解析错误")
				return
			}
			err = global.TdDb.InsertEvent(data)
			if err != nil {
				global.Log.Error("连接事件数据添加错误", err)
			}
		}
	}()
}

// 获取规则实体
func getRuleChainInstance(etoken *model.DeviceAuth) *rule_engine.RuleChainInstance {
	defer func() {
		if err := recover(); err != nil {
			global.Log.Error(err)
		}
	}()

	key := etoken.ProductId
	instance, err := cache.ComputeIfAbsentProductRule(key, func(k any) (any, error) {
		one, err := services.ProductModelDao.FindOne(k.(string))
		if err != nil {
			return nil, err
		}
		rule := ruleService.RuleChainModelDao.FindOne(one.RuleChainId)
		var lfData ruleEntity.RuleDataJson
		err = tool.StringToStruct(rule.RuleDataJson, &lfData)
		if err != nil {
			return nil, err
		}
		code, _ := json.Marshal(lfData.LfData.DataCode)
		//新建规则链实体
		instance, errs := rule_engine.NewRuleChainInstance(rule.Id, code)
		if errs != nil {
			global.Log.Error("规则链初始化失败", errs)
			return nil, errs
		}
		return instance, nil
	})
	biz.ErrIsNil(err, "缓存读取规则链失败")
	if ruleData, ok := instance.(*rule_engine.RuleChainInstance); ok {
		return ruleData
	}
	return nil
}

// 构建规则链执行的消息
func buildRuleMessage(etoken *model.DeviceAuth, msgVals map[string]interface{}, msgType string) *message.Message {
	metadataVals := map[string]interface{}{
		"deviceId":       etoken.DeviceId,
		"deviceName":     etoken.Name,
		"deviceType":     etoken.DeviceType,
		"deviceProtocol": etoken.DeviceProtocol,
		"productId":      etoken.ProductId,
		"orgId":          etoken.OrgId,
		"owner":          etoken.Owner,
	}
	return message.NewMessage(etoken.Owner, msgType, msgVals, metadataVals)
}

func SendZtWebsocket(deviceId, message string) {
	msgVals := make(map[string]interface{})
	if err := json.Unmarshal([]byte(message), &msgVals); err != nil {
		return
	}
	twinData := map[string]interface{}{
		"twinId": deviceId,
		"attrs":  msgVals,
	}
	data, _ := json.Marshal(twinData)
	for stageid := range websocket.Wsp {
		CJNR := fmt.Sprintf(`{"MESSAGETYPE":"01","MESSAGECONTENT": %s}`, string(data))
		websocket.SendMessage(CJNR, stageid)
	}
}

// SetDeviceShadow 设置设备点
func SetDeviceShadow(etoken *model.DeviceAuth, msgVals map[string]interface{}, msgType string) {
	defer func() {
		if err := recover(); err != nil {
			global.Log.Error(err)
		}
	}()

	if msgType == message.RowMes {
		msgType = message.TelemetryMes
	}
	for key, value := range msgVals {
		if message.AttributesMes == msgType {
			err := shadow.DeviceShadowInstance.SetDevicePoint(etoken.Name, global.TslAttributesType, key, value)
			biz.ErrIsNilAppendErr(err, "设置设备影子点失败")
		}
		if message.TelemetryMes == msgType {
			err := shadow.DeviceShadowInstance.SetDevicePoint(etoken.Name, global.TslTelemetryType, key, value)
			biz.ErrIsNilAppendErr(err, "设置设备影子点失败")
		}
	}
}
