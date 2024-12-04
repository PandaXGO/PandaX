package initialize

import (
	"encoding/json"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	"pandax/pkg/events"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine"
	"pandax/pkg/tool"
)

// 初始化事件监听
func InitEvents() {
	// 监听规则链改变 更新所有绑定改规则链的产品
	global.EventEmitter.On(events.ProductChainRuleEvent, func(ruleId string, codeData string) {
		global.Log.Infof("规则链%s变更", ruleId)
		list, _ := services.ProductModelDao.FindListByRule(ruleId)
		if list != nil {
			var lfData ruleEntity.RuleDataJson
			err := tool.StringToStruct(codeData, &lfData)
			if err != nil {
				global.Log.Error("规则链序列化失败", err)
				return
			}
			code, err := json.Marshal(lfData.DataCode)
			if err != nil {
				global.Log.Error("规则链序列化失败", err)
				return
			}
			//新建规则链实体
			instance, err := rule_engine.NewRuleChainInstance(ruleId, code)
			if err != nil {
				global.Log.Error("规则链初始化失败", err)
				return
			}
			for _, product := range *list {
				rule_engine.RuleEngine.SaveRuleInstance(product.Id, instance)
			}
		}
	})
}
