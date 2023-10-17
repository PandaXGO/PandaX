package hook_message_work

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PandaXGO/PandaKit/biz"
	"log"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	ruleService "pandax/apps/rule/services"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/global_model"
	"pandax/pkg/rule_engine"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/shadow"
	"pandax/pkg/tool"
	"pandax/pkg/websocket"
	"strings"
)

// 消息处理模块
func (s *HookService) MessageWork() {
	for {
		select {
		case msg := <-s.MessageCh:
			s.handleOne(msg) // 处理消息
		}
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
			// 获取规则链代码
			chain := getRuleChain(msg.DeviceAuth)
			if chain == nil {
				return
			}
			dataCode := chain.LfData.DataCode
			code, err := json.Marshal(dataCode)
			//新建规则链实体
			instance, errs := rule_engine.NewRuleChainInstance(code)
			if len(errs) > 0 {
				global.Log.Error("规则链初始化失败", errs[0])
				return
			}
			ruleMessage := buildRuleMessage(msg.DeviceAuth, msgVals, msg.Type)
			err = instance.StartRuleChain(context.Background(), ruleMessage)
			if err != nil {
				global.Log.Error("规则链执行失败", errs)
			}
			// 保存设备影子
			if msg.Type != message.RpcRequestFromDevice {
				SetDeviceShadow(msg.DeviceAuth, ruleMessage.Msg, msg.Type)
			}
		case message.DisConnectMes, message.ConnectMes:
			//检测设备影子并修改设备影子状态
			if msg.Type == message.ConnectMes {
				shadow.InitDeviceShadow(msg.DeviceAuth.Name, msg.DeviceAuth.ProductId)
				err := shadow.DeviceShadowInstance.SetOnline(msg.DeviceAuth.Name)
				log.Println(err)
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

func getRuleChain(etoken *global_model.DeviceAuth) *ruleEntity.RuleDataJson {
	defer func() {
		if err := recover(); err != nil {
			global.Log.Error(err)
		}
	}()
	key := etoken.ProductId
	get, err := global.ProductCache.ComputeIfAbsent(key, func(k any) (any, error) {
		one := services.ProductModelDao.FindOne(k.(string))
		rule := ruleService.RuleChainModelDao.FindOne(one.RuleChainId)
		return rule.RuleDataJson, nil
	})
	biz.ErrIsNil(err, "缓存读取规则链失败")
	ruleData := ruleEntity.RuleDataJson{}
	err = tool.StringToStruct(get.(string), &ruleData)
	biz.ErrIsNil(err, "规则链数据转化失败")
	return &ruleData
}

func buildRuleMessage(etoken *global_model.DeviceAuth, msgVals map[string]interface{}, msgType string) *message.Message {
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
	for stageid, _ := range websocket.Wsp {
		CJNR := fmt.Sprintf(`{"MESSAGETYPE":"01","MESSAGECONTENT": %s}`, string(data))
		websocket.SendMessage(CJNR, stageid)
	}
}

// SetDeviceShadow 设置设备点
func SetDeviceShadow(etoken *global_model.DeviceAuth, msgVals map[string]interface{}, msgType string) {
	defer func() {
		if err := recover(); &err != nil {
			global.Log.Error(err)
		}
	}()

	if msgType == message.RowMes {
		msgType = message.TelemetryMes
	}
	template := services.ProductTemplateModelDao.FindList(entity.ProductTemplate{Classify: strings.ToLower(msgType), Pid: etoken.ProductId})
	for _, tel := range *template {
		if _, ok := msgVals[tel.Key]; !ok {
			continue
		}
		if message.AttributesMes == msgType {
			err := shadow.DeviceShadowInstance.SetDevicePoint(etoken.Name, global.TslAttributesType, tel.Key, msgVals[tel.Key])
			biz.ErrIsNilAppendErr(err, "设置设备影子点失败")
		}
		if message.TelemetryMes == msgType {
			log.Println(etoken.Name)
			err := shadow.DeviceShadowInstance.SetDevicePoint(etoken.Name, global.TslTelemetryType, tel.Key, msgVals[tel.Key])
			biz.ErrIsNilAppendErr(err, "设置设备影子点失败")
		}
	}
}
