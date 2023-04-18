// ==========================================================================
// 生成日期：2023-03-29 20:01:11 +0800 CST
// 生成路径: apps/visual/services/rulechain.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualRuleChainMsgLogModel interface {
		Insert(data entity.VisualRuleChainMsgLog) *entity.VisualRuleChainMsgLog
		FindListPage(page, pageSize int, data entity.VisualRuleChainMsgLog) (*[]entity.VisualRuleChainMsgLog, int64)
		Delete(ids []string)
	}

	ruleChainLogModelImpl struct {
		table string
	}
)

var VisualRuleChainMsgLogModelDao VisualRuleChainMsgLogModel = &ruleChainLogModelImpl{
	table: `visual_rule_chain_msg_log`,
}

func (m *ruleChainLogModelImpl) Insert(data entity.VisualRuleChainMsgLog) *entity.VisualRuleChainMsgLog {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加规则链失败")
	return &data
}

func (m *ruleChainLogModelImpl) FindListPage(page, pageSize int, data entity.VisualRuleChainMsgLog) (*[]entity.VisualRuleChainMsgLog, int64) {
	list := make([]entity.VisualRuleChainMsgLog, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	db.Where("delete_time IS NULL")
	if data.DeviceName != "" {
		db = db.Where("device_name = ?", data.DeviceName)
	}
	if data.MessageId != "" {
		db = db.Where("message_id = ?", data.MessageId)
	}
	if data.MsgType != "" {
		db = db.Where("msg_type = ?", data.MsgType)
	}
	err := db.Count(&total).Error
	err = db.Order("create_at").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询规则链分页列表失败")
	return &list, total
}

func (m *ruleChainLogModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualRuleChainMsgLog{}, "message_id in (?)", ids).Error, "删除规则链失败")
}
