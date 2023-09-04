package iothub

import (
	"context"
	"encoding/json"
	"fmt"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/tool"
	"pandax/pkg/websocket"
)

// 消息处理模块
func (s *HookService) MessageWork() {
	for {
		select {
		case msg := <-s.messageCh:
			s.handleOne(msg) // 处理消息
		}
	}
}

func (s *HookService) handleOne(msg *DeviceEventInfo) {
	if s.ch != nil { // 用于并发限制
		s.ch <- struct{}{}
	}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		etoken := &tool.DeviceAuth{}
		err := global.RedisDb.Get(msg.DeviceId, etoken)
		if err != nil {
			return
		}
		switch msg.Type {
		case message.RowMes, message.AttributesMes, message.TelemetryMes:
			// 发送websocket到云组态
			if msg.Type == message.TelemetryMes {
				go SendZtWebsocket(msg.DeviceId, msg.Datas)
			}
			// 业务逻辑执行
			// 获取规则链代码
			chain := getRuleChain(etoken)
			dataCode := chain.LfData.DataCode
			code, err := json.Marshal(dataCode)
			//新建规则链实体
			instance, errs := rule_engine.NewRuleChainInstance(code)
			if len(errs) > 0 {
				global.Log.Error("规则链初始化失败", errs[0])
				return
			}
			ruleMessage := buildRuleMessage(etoken, msg, msg.Type)
			err = instance.StartRuleChain(context.Background(), ruleMessage)
			if err != nil {
				global.Log.Error("规则链执行失败", errs)
			}
		// Rpc请求
		case message.RpcRequestMes:
			chain := getRuleChain(etoken)
			dataCode := chain.LfData.DataCode
			code, err := json.Marshal(dataCode)
			//新建规则链实体
			instance, errs := rule_engine.NewRuleChainInstance(code)
			if len(errs) > 0 {
				global.Log.Error("规则链初始化失败", errs[0])
				return
			}
			ruleMessage := buildRuleMessage(etoken, msg, msg.Type)
			err = instance.StartRuleChain(context.Background(), ruleMessage)
			if err != nil {
				global.Log.Error("规则链执行失败", errs)
			}
		case message.DisConnectMes, message.ConnectMes:
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

func getRuleChain(etoken *tool.DeviceAuth) *ruleEntity.RuleDataJson {
	ruleData := ruleEntity.RuleDataJson{}
	err := global.RedisDb.Get(etoken.ProductId, &ruleData)
	if err != nil {
		return nil
	}
	return &ruleData
}

func buildRuleMessage(etoken *tool.DeviceAuth, dei *DeviceEventInfo, msgType string) *message.Message {
	metadataVals := map[string]interface{}{
		"deviceId":   etoken.DeviceId,
		"deviceName": etoken.Name,
		"deviceType": etoken.DeviceType,
		"productId":  etoken.ProductId,
	}
	msgVals := make(map[string]interface{})
	json.Unmarshal([]byte(dei.Datas), &msgVals)
	return message.NewMessage(etoken.User, msgType, msgVals, metadataVals)
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
