// ==========================================================================
// 生成日期：2023-03-29 20:01:11 +0800 CST
// 生成路径: apps/visual/services/rulechain.go
// 生成人：panda
// ==========================================================================

package services

import (
	"errors"
	"pandax/apps/rule/entity"
	"pandax/pkg/events"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
)

type (
	RuleChainModel interface {
		Insert(data entity.RuleChain) (*entity.RuleChain, error)
		FindOne(id string) (*entity.RuleChain, error)
		FindOneByRoot() (*entity.RuleChain, error)
		UpdateByRoot() error
		FindListPage(page, pageSize int, data entity.RuleChain) (*[]entity.RuleChainBase, int64, error)
		FindList(data entity.RuleChain) (*[]entity.RuleChain, error)
		FindListBaseLabel(data entity.RuleChain) (*[]entity.RuleChainBaseLabel, error)
		Update(data entity.RuleChain) (*entity.RuleChain, error)
		Delete(ids []string) error
	}

	ruleChainModelImpl struct {
		table string
	}
)

var RuleChainModelDao RuleChainModel = &ruleChainModelImpl{
	table: `rule_chain`,
}

func (m *ruleChainModelImpl) Insert(data entity.RuleChain) (*entity.RuleChain, error) {
	tx := global.Db.Begin()
	// 如果新增的链为主链，那么将原来的设置为普通连
	if data.Root == "1" {
		err := m.UpdateByRoot()
		if err != nil {
			return nil, errors.New("修改主链错误")
		}
	}
	err := global.Db.Table(m.table).Create(&data).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &data, nil
}

func (m *ruleChainModelImpl) FindOne(id string) (*entity.RuleChain, error) {
	resData := new(entity.RuleChain)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	return resData, err
}

func (m *ruleChainModelImpl) FindOneByRoot() (*entity.RuleChain, error) {
	resData := new(entity.RuleChain)
	db := global.Db.Table(m.table).Where("root = ?", 1)
	err := db.First(resData).Error
	return resData, err
}

// UpdateByRoot 修改主链为普通链
func (m *ruleChainModelImpl) UpdateByRoot() error {
	tx := global.Db.Table(m.table).Where("root = ?", "1").Update("root", "0")
	return tx.Error
}

func (m *ruleChainModelImpl) FindListPage(page, pageSize int, data entity.RuleChain) (*[]entity.RuleChainBase, int64, error) {
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
	err := model.OrgAuthSet(db, data.RoleId, data.Owner)
	if err != nil {
		return &list, total, err
	}
	err = db.Count(&total).Error
	if err != nil {
		return &list, total, err
	}
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *ruleChainModelImpl) FindList(data entity.RuleChain) (*[]entity.RuleChain, error) {
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
	err := model.OrgAuthSet(db, data.RoleId, data.Owner)
	if err != nil {
		return &list, err
	}
	err = db.Order("create_time").Find(&list).Error
	return &list, err
}

func (m *ruleChainModelImpl) FindListBaseLabel(data entity.RuleChain) (*[]entity.RuleChainBaseLabel, error) {
	list := make([]entity.RuleChainBaseLabel, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.RuleName != "" {
		db = db.Where("rule_name like ?", "%"+data.RuleName+"%")
	}
	// 组织数据访问权限
	err := model.OrgAuthSet(db, data.RoleId, data.Owner)
	if err != nil {
		return &list, err
	}
	err = db.Find(&list).Error
	return &list, err
}

func (m *ruleChainModelImpl) Update(data entity.RuleChain) (*entity.RuleChain, error) {
	tx := global.Db.Begin()
	// 如果新增的链为主链，那么将原来的设置为普通连
	if data.Root == "1" {
		err := m.UpdateByRoot()
		if err != nil {
			return nil, err
		}
	}
	err := global.Db.Table(m.table).Updates(&data).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	//更改本地规则链缓存
	if data.RuleDataJson != "" {
		go global.EventEmitter.Emit(events.ProductChainRuleEvent, data.Id, data.RuleDataJson)
	}
	return &data, nil
}

func (m *ruleChainModelImpl) Delete(ids []string) error {
	return global.Db.Table(m.table).Delete(&entity.RuleChain{}, "id in (?)", ids).Error
}
