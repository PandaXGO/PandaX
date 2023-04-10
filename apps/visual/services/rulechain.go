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
	VisualRuleChainModel interface {
		Insert(data entity.VisualRuleChain) *entity.VisualRuleChain
		FindOne(id string) *entity.VisualRuleChain
		FindListPage(page, pageSize int, data entity.VisualRuleChain) (*[]entity.VisualRuleChain, int64)
		FindList(data entity.VisualRuleChain) *[]entity.VisualRuleChain
		Update(data entity.VisualRuleChain) *entity.VisualRuleChain
		Delete(ids []string)
	}

	ruleChainModelImpl struct {
		table string
	}
)

var VisualRuleChainModelDao VisualRuleChainModel = &ruleChainModelImpl{
	table: `visual_rule_chain`,
}

func (m *ruleChainModelImpl) Insert(data entity.VisualRuleChain) *entity.VisualRuleChain {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加规则链失败")
	return &data
}

func (m *ruleChainModelImpl) FindOne(id string) *entity.VisualRuleChain {
	resData := new(entity.VisualRuleChain)
	db := global.Db.Table(m.table).Where("rule_id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询规则链失败")
	return resData
}

func (m *ruleChainModelImpl) FindListPage(page, pageSize int, data entity.VisualRuleChain) (*[]entity.VisualRuleChain, int64) {
	list := make([]entity.VisualRuleChain, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	db.Where("delete_time IS NULL")
	if data.UserId != "" {
		db = db.Where("user_id = ?", data.UserId)
	}
	if data.RuleName != "" {
		db = db.Where("rule_name like ?", "%"+data.RuleName+"%")
	}
	if data.RuleRemark != "" {
		db = db.Where("rule_remark like ?", "%"+data.RuleRemark+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询规则链分页列表失败")
	return &list, total
}

func (m *ruleChainModelImpl) FindList(data entity.VisualRuleChain) *[]entity.VisualRuleChain {
	list := make([]entity.VisualRuleChain, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	db.Where("delete_time IS NULL")
	if data.UserId != "" {
		db = db.Where("user_id = ?", data.UserId)
	}
	if data.RuleName != "" {
		db = db.Where("rule_name like ?", "%"+data.RuleName+"%")
	}
	if data.RuleRemark != "" {
		db = db.Where("rule_remark like ?", "%"+data.RuleRemark+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询规则链列表失败")
	return &list
}

func (m *ruleChainModelImpl) Update(data entity.VisualRuleChain) *entity.VisualRuleChain {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改规则链失败")
	return &data
}

func (m *ruleChainModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualRuleChain{}, "rule_id in (?)", ids).Error, "删除规则链失败")
}
