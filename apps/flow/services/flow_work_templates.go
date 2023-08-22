// ==========================================================================
// 生成日期：2023-03-29 19:46:55 +0800 CST
// 生成路径: apps/flow/services/flow_work_templates.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/flow/entity"
	"pandax/pkg/global"
)

type (
	FlowWorkTemplatesModel interface {
		Insert(data entity.FlowWorkTemplates) *entity.FlowWorkTemplates
		FindOne(id int64) *entity.FlowWorkTemplates
		FindListPage(page, pageSize int, data entity.FlowWorkTemplates) (*[]entity.FlowWorkTemplates, int64)
		FindList(data entity.FlowWorkTemplates) *[]entity.FlowWorkTemplates
		Update(data entity.FlowWorkTemplates) *entity.FlowWorkTemplates
		Delete(ids []int64)
	}

	worktemplatesModelImpl struct {
		table string
	}
)

var FlowWorkTemplatesModelDao FlowWorkTemplatesModel = &worktemplatesModelImpl{
	table: `flow_work_templates`,
}

func (m *worktemplatesModelImpl) Insert(data entity.FlowWorkTemplates) *entity.FlowWorkTemplates {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加工作流模板失败")
	return &data
}

func (m *worktemplatesModelImpl) FindOne(id int64) *entity.FlowWorkTemplates {
	resData := new(entity.FlowWorkTemplates)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询工作流模板失败")
	return resData
}

func (m *worktemplatesModelImpl) FindListPage(page, pageSize int, data entity.FlowWorkTemplates) (*[]entity.FlowWorkTemplates, int64) {
	list := make([]entity.FlowWorkTemplates, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Remarks != "" {
		db = db.Where("remarks = ?", data.Remarks)
	}
	db.Where("delete_time IS NULL")
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Creator != 0 {
		db = db.Where("creator = ?", data.Creator)
	}
	if data.Remarks != "" {
		db = db.Where("name like ?", "%"+data.Remarks+"%")
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询工作流模板分页列表失败")
	return &list, total
}

func (m *worktemplatesModelImpl) FindList(data entity.FlowWorkTemplates) *[]entity.FlowWorkTemplates {
	list := make([]entity.FlowWorkTemplates, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Remarks != "" {
		db = db.Where("remarks = ?", data.Remarks)
	}
	db.Where("delete_time IS NULL")
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Creator != 0 {
		db = db.Where("creator = ?", data.Creator)
	}
	if data.Remarks != "" {
		db = db.Where("name like ?", "%"+data.Remarks+"%")
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询工作流模板列表失败")
	return &list
}

func (m *worktemplatesModelImpl) Update(data entity.FlowWorkTemplates) *entity.FlowWorkTemplates {
	err := global.Db.Table(m.table).Updates(&data).Error
	if err != nil {
		global.Log.Error(err)
		biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改工作流模板失败")
	}

	return &data
}

func (m *worktemplatesModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.FlowWorkTemplates{}, "id in (?)", ids).Error, "删除工作流模板失败")
}
