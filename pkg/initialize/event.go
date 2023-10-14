package initialize

import (
	"log"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/pkg/events"
	"pandax/pkg/global"
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
				log.Println("修改了产品的规则链", product.Id)
				global.ProductCache.Put(product.Id, codeData)
			}
		}
	})
}
