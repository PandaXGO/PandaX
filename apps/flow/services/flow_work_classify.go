// ==========================================================================
// 生成日期：2022-08-24 22:02:33 +0800 CST
// 生成路径: apps/flow/services/flow_work_classify.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/flow/entity"
	"pandax/pkg/global"
)

type (
	FlowWorkClassifyModel interface {
		Insert(data entity.FlowWorkClassify) *entity.FlowWorkClassify
		FindOne(id int64) *entity.FlowWorkClassify
		FindListPage(page, pageSize int, data entity.FlowWorkClassify) (*[]entity.FlowWorkClassify, int64)
		FindList(data entity.FlowWorkClassify) *[]entity.FlowWorkClassify
		Update(data entity.FlowWorkClassify) *entity.FlowWorkClassify
		Delete(ids []int64)
	}

	workclassifyModelImpl struct {
		table string
	}
)

var FlowWorkClassifyModelDao FlowWorkClassifyModel = &workclassifyModelImpl{
	table: `flow_work_classify`,
}

func (m *workclassifyModelImpl) Insert(data entity.FlowWorkClassify) *entity.FlowWorkClassify {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加工作流分类失败")
	return &data
}

func (m *workclassifyModelImpl) FindOne(id int64) *entity.FlowWorkClassify {
	resData := new(entity.FlowWorkClassify)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询工作流分类失败")
	return resData
}

func (m *workclassifyModelImpl) FindListPage(page, pageSize int, data entity.FlowWorkClassify) (*[]entity.FlowWorkClassify, int64) {
	list := make([]entity.FlowWorkClassify, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Creator != 0 {
		db = db.Where("creator = ?", data.Creator)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询工作流分类分页列表失败")
	return &list, total
}

func (m *workclassifyModelImpl) FindList(data entity.FlowWorkClassify) *[]entity.FlowWorkClassify {
	list := make([]entity.FlowWorkClassify, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Creator != 0 {
		db = db.Where("creator = ?", data.Creator)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询工作流分类列表失败")
	return &list
}

func (m *workclassifyModelImpl) Update(data entity.FlowWorkClassify) *entity.FlowWorkClassify {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改工作流分类失败")
	return &data
}

func (m *workclassifyModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.FlowWorkClassify{}, "id in (?)", ids).Error, "删除工作流分类失败")
}
