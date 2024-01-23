// ==========================================================================
// 生成日期：2023-03-29 20:01:11 +0800 CST
// 生成路径: apps/visual/services/rulechain.go
// 生成人：panda
// ==========================================================================

package services

import (
	"pandax/apps/rule/entity"
	"pandax/kit/biz"
	"pandax/pkg/events"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
)

type (
	RuleChainModel interface {
		Insert(data entity.RuleChain) *entity.RuleChain
		FindOne(id string) *entity.RuleChain
		FindOneByRoot() *entity.RuleChain
		UpdateByRoot() error
		FindListPage(page, pageSize int, data entity.RuleChain) (*[]entity.RuleChainBase, int64)
		FindList(data entity.RuleChain) *[]entity.RuleChain
		FindListBaseLabel(data entity.RuleChain) *[]entity.RuleChainBaseLabel
		Update(data entity.RuleChain) *entity.RuleChain
		Delete(ids []string)
	}

	ruleChainModelImpl struct {
		table string
	}
)

var RuleChainModelDao RuleChainModel = &ruleChainModelImpl{
	table: `rule_chain`,
}

func (m *ruleChainModelImpl) Insert(data entity.RuleChain) *entity.RuleChain {
	tx := global.Db.Begin()
	// 如果新增的链为主链，那么将原来的设置为普通连
	if data.Root == "1" {
		err := m.UpdateByRoot()
		biz.ErrIsNil(err, "修改主链错误")
	}
	err := global.Db.Table(m.table).Create(&data).Error
	if err != nil {
		tx.Rollback()
	}
	biz.ErrIsNil(err, "添加规则链失败")
	tx.Commit()
	return &data
}

func (m *ruleChainModelImpl) FindOne(id string) *entity.RuleChain {
	resData := new(entity.RuleChain)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error

	biz.ErrIsNil(err, "查询规则链失败")
	return resData
}

func (m *ruleChainModelImpl) FindOneByRoot() *entity.RuleChain {
	resData := new(entity.RuleChain)
	db := global.Db.Table(m.table).Where("root = ?", 1)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询规则链失败")
	return resData
}

// UpdateByRoot 修改主链为普通链
func (m *ruleChainModelImpl) UpdateByRoot() error {
	tx := global.Db.Table(m.table).Where("root = ?", "1").Update("root", "0")
	return tx.Error
}

func (m *ruleChainModelImpl) FindListPage(page, pageSize int, data entity.RuleChain) (*[]entity.RuleChainBase, int64) {
	list := make([]entity.RuleChainBase, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.RuleName != "" {
		db = db.Where("rule_name like ?", "%"+data.RuleName+"%")
	}
	if data.RuleRemark != "" {
		db = db.Where("rule_remark like ?", "%"+data.RuleRemark+"%")
	}
	// 组织数据访问权限
	model.OrgAuthSet(db, data.RoleId, data.Owner)
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询规则链分页列表失败")
	return &list, total
}

func (m *ruleChainModelImpl) FindList(data entity.RuleChain) *[]entity.RuleChain {
	list := make([]entity.RuleChain, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.RuleName != "" {
		db = db.Where("rule_name like ?", "%"+data.RuleName+"%")
	}
	if data.RuleRemark != "" {
		db = db.Where("rule_remark like ?", "%"+data.RuleRemark+"%")
	}
	// 组织数据访问权限
	model.OrgAuthSet(db, data.RoleId, data.Owner)
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询规则链列表失败")
	return &list
}

func (m *ruleChainModelImpl) FindListBaseLabel(data entity.RuleChain) *[]entity.RuleChainBaseLabel {
	list := make([]entity.RuleChainBaseLabel, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.RuleName != "" {
		db = db.Where("rule_name like ?", "%"+data.RuleName+"%")
	}
	// 组织数据访问权限
	model.OrgAuthSet(db, data.RoleId, data.Owner)
	biz.ErrIsNil(db.Find(&list).Error, "查询规则链列表失败")
	return &list
}

func (m *ruleChainModelImpl) Update(data entity.RuleChain) *entity.RuleChain {
	tx := global.Db.Begin()
	// 如果新增的链为主链，那么将原来的设置为普通连
	if data.Root == "1" {
		err := m.UpdateByRoot()
		biz.ErrIsNil(err, "修改主链错误")
	}
	err := global.Db.Table(m.table).Updates(&data).Error
	if err != nil {
		tx.Rollback()
	}
	biz.ErrIsNil(err, "修改规则链失败")
	tx.Commit()
	//更改本地规则链缓存
	if data.RuleDataJson != "" {
		go global.EventEmitter.Emit(events.ProductChainRuleEvent, data.Id, data.RuleDataJson)
	}
	return &data
}

func (m *ruleChainModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.RuleChain{}, "id in (?)", ids).Error, "删除规则链失败")
}
