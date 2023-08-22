package initialize

import (
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/pkg/events"
	"pandax/pkg/global"
	"time"
)

// 初始化事件监听
func InitEvents() {
	// 监听规则链改变 更新所有绑定改规则链的产品
	global.EventEmitter.On(events.ProductChainRuleEvent, func(ruleId, codeData string) {
		global.Log.Infof("规则链%s变更", ruleId)
		list := services.ProductModelDao.FindList(entity.Product{
			RuleChainId: ruleId,
		})
		if list != nil {
			for _, product := range *list {
				err := global.RedisDb.Set(product.Id, codeData, time.Hour*24*365)
				if err != nil {
					global.Log.Errorf("事件监听执行错误：%s", err.Error())
				}
			}
		}
	})
}
