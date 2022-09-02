// ==========================================================================
// 生成日期：2022-09-02 15:49:39 +0800 CST
// 生成路径: apps/rule/services/rule_notice.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/rule/entity"
	"pandax/pkg/global"
)

type (
	RuleNoticeModel interface {
		Insert(data entity.RuleNotice) *entity.RuleNotice
		FindOne(id int64) *entity.RuleNotice
		FindListPage(page, pageSize int, data entity.RuleNotice) (*[]entity.RuleNotice, int64)
		FindList(data entity.RuleNotice) *[]entity.RuleNotice
		Update(data entity.RuleNotice) *entity.RuleNotice
		Delete(ids []int64)
	}

	noticeModelImpl struct {
		table string
	}
)

var RuleNoticeModelDao RuleNoticeModel = &noticeModelImpl{
	table: `rule_notice`,
}

func (m *noticeModelImpl) Insert(data entity.RuleNotice) *entity.RuleNotice {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加规则通知配置失败")
	return &data
}

func (m *noticeModelImpl) FindOne(id int64) *entity.RuleNotice {
	resData := new(entity.RuleNotice)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询规则通知配置失败")
	return resData
}

func (m *noticeModelImpl) FindListPage(page, pageSize int, data entity.RuleNotice) (*[]entity.RuleNotice, int64) {
	list := make([]entity.RuleNotice, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Description != "" {
		db = db.Where("description = ?", data.Description)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Category != "" {
		db = db.Where("category = ?", data.Category)
	}
	if data.Model != "" {
		db = db.Where("model = ?", data.Model)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询规则通知配置分页列表失败")
	return &list, total
}

func (m *noticeModelImpl) FindList(data entity.RuleNotice) *[]entity.RuleNotice {
	list := make([]entity.RuleNotice, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Description != "" {
		db = db.Where("description = ?", data.Description)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Category != "" {
		db = db.Where("category = ?", data.Category)
	}
	if data.Model != "" {
		db = db.Where("model = ?", data.Model)
	}
	db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询规则通知配置列表失败")
	return &list
}

func (m *noticeModelImpl) Update(data entity.RuleNotice) *entity.RuleNotice {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改规则通知配置失败")
	return &data
}

func (m *noticeModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.RuleNotice{}, "id in (?)", ids).Error, "删除规则通知配置失败")
}
