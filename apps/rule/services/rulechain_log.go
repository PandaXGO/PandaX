// ==========================================================================
// 生成日期：2023-03-29 20:01:11 +0800 CST
// 生成路径: apps/visual/services/rulechain.go
// 生成人：panda
// ==========================================================================

package services

import (
	"pandax/apps/rule/entity"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
)

type (
	RuleChainMsgLogModel interface {
		Insert(data entity.RuleChainMsgLog) (*entity.RuleChainMsgLog, error)
		FindListPage(page, pageSize int, data entity.RuleChainMsgLog) (*[]entity.RuleChainMsgLog, int64, error)
		Delete(data entity.RuleChainMsgLog) error
	}

	ruleChainLogModelImpl struct {
		table string
	}
)

var RuleChainMsgLogModelDao RuleChainMsgLogModel = &ruleChainLogModelImpl{
	table: `rule_chain_msg_log`,
}

func (m *ruleChainLogModelImpl) Insert(data entity.RuleChainMsgLog) (*entity.RuleChainMsgLog, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *ruleChainLogModelImpl) FindListPage(page, pageSize int, data entity.RuleChainMsgLog) (*[]entity.RuleChainMsgLog, int64, error) {
	list := make([]entity.RuleChainMsgLog, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DeviceName != "" {
		db = db.Where("device_name = ?", data.DeviceName)
	}
	if data.MessageId != "" {
		db = db.Where("message_id = ?", data.MessageId)
	}
	if data.MsgType != "" {
		db = db.Where("msg_type = ?", data.MsgType)
	}
	// 组织数据访问权限
	err := model.OrgAuthSet(db, data.RoleId, data.Owner)
	if err != nil {
		return nil, 0, err
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("create_at").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *ruleChainLogModelImpl) Delete(data entity.RuleChainMsgLog) error {
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DeviceName != "" {
		db = db.Where("device_name = ?", data.DeviceName)
	}
	if data.MessageId != "" {
		db = db.Where("message_id = ?", data.MessageId)
	}
	if data.MsgType != "" {
		db = db.Where("msg_type = ?", data.MsgType)
	}
	return db.Delete(&entity.RuleChainMsgLog{}).Error
}
