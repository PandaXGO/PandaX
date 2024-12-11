package hook_message_work

import (
	"encoding/json"
	"fmt"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	ruleService "pandax/apps/rule/services"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"pandax/pkg/rule_engine"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/shadow"
	"pandax/pkg/tdengine"
	"pandax/pkg/tool"
	"pandax/pkg/websocket"
	"time"

	"github.com/kakuilan/kgo"
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
		case message.RowMes, message.AttributesMes, message.TelemetryMes, message.RpcRequestFromDevice, message.UpEventMes:
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
			instance, err := getRuleChainInstance(msg.DeviceAuth)
			if err != nil {
				global.Log.Error("获取设备实体失败", err)
				return
			}
			ruleMessage := buildRuleMessage(msg.DeviceAuth, msgVals, msg.Type)
			err = rule_engine.RuleEngine.StartRuleInstance(instance, ruleMessage)
			if err != nil {
				global.Log.Error("规则链执行失败", err)
				return
			}
			// 保存事件信息
			if msg.Type == message.UpEventMes {
				tsl, err := services.ProductTemplateModelDao.FindOneByKey(msg.DeviceId, msg.Identifier)
				if err != nil {
					return
				}
				ci := &tdengine.Events{
					DeviceId: msg.DeviceId,
					Name:     msg.Identifier,
					Type:     tsl.Type,
					Content:  msg.Datas,
					Ts:       time.Now().Format("2006-01-02 15:04:05.000"),
				}
				data, err := kgo.KConv.Struct2Map(ci, "")
				if err != nil {
					global.Log.Error("事件格式转化错误")
					return
				}
				err = global.TdDb.InsertEvent(data)
				if err != nil {
					global.Log.Error("事件添加错误", err)
				}
			}
			// 刷新设备状态
			shadow.DeviceShadowInstance.RefreshDeviceStatus(msg.DeviceAuth.Name)
		case message.DisConnectMes, message.ConnectMes:
			// 更改设备在线状态
			isHas := shadow.DeviceShadowInstance.HasDevice(msg.DeviceAuth.Name)
			if !isHas {
				shadow.InitDeviceShadow(msg.DeviceAuth.Name)
			}
			if msg.Type == message.ConnectMes {
				shadow.DeviceShadowInstance.SetOnline(msg.DeviceAuth.Name)
			} else {
				shadow.DeviceShadowInstance.SetOffline(msg.DeviceAuth.Name)
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
func getRuleChainInstance(etoken *model.DeviceAuth) (*rule_engine.RuleChainInstance, error) {
	key := etoken.ProductId
	instance := rule_engine.RuleEngine.GetRuleInstance(key)
	if instance == nil {
		one, err := services.ProductModelDao.FindOne(key)
		if err != nil {
			return nil, err
		}
		rule, err := ruleService.RuleChainModelDao.FindOne(one.RuleChainId)
		if err != nil {
			return nil, err
		}
		var lfData ruleEntity.RuleDataJson
		err = tool.StringToStruct(rule.RuleDataJson, &lfData)
		if err != nil {
			return nil, err
		}
		code, _ := json.Marshal(lfData.DataCode)
		//新建规则链实体
		instance, err = rule_engine.NewRuleChainInstance(rule.Id, code)
		if err != nil {
			return nil, err
		}
		_, err = rule_engine.RuleEngine.SaveRuleInstance(key, instance)
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// 构建规则链执行的消息
func buildRuleMessage(etoken *model.DeviceAuth, msgVals map[string]interface{}, msgType string) *message.Message {
	metadataVals := map[string]interface{}{
		"deviceId":       etoken.DeviceId,
		"deviceName":     etoken.Name,
		"deviceType":     etoken.DeviceType,
		"deviceProtocol": etoken.Protocol,
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
