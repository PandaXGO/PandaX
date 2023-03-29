// ==========================================================================
// 生成日期：2023-03-29 20:01:11 +0800 CST
// 生成路径: apps/flow/services/flow_work_info.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/flow/entity"
	"pandax/pkg/global"
)

type (
	FlowWorkInfoModel interface {
		Insert(data entity.FlowWorkInfo) *entity.FlowWorkInfo
		FindOne(id int64) *entity.FlowWorkInfo
		FindListPage(page, pageSize int, data entity.FlowWorkInfo) (*[]entity.FlowWorkInfo, int64)
		FindList(data entity.FlowWorkInfo) *[]entity.FlowWorkInfo
		Update(data entity.FlowWorkInfo) *entity.FlowWorkInfo
		Delete(ids []int64)
	}

	workinfoModelImpl struct {
		table string
	}
)

var FlowWorkInfoModelDao FlowWorkInfoModel = &workinfoModelImpl{
	table: `flow_work_info`,
}

func (m *workinfoModelImpl) Insert(data entity.FlowWorkInfo) *entity.FlowWorkInfo {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加流程失败")
	return &data
}

func (m *workinfoModelImpl) FindOne(id int64) *entity.FlowWorkInfo {
	resData := new(entity.FlowWorkInfo)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询流程失败")
	return resData
}

func (m *workinfoModelImpl) FindListPage(page, pageSize int, data entity.FlowWorkInfo) (*[]entity.FlowWorkInfo, int64) {
	list := make([]entity.FlowWorkInfo, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	db.Where("delete_time IS NULL")
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Creator != 0 {
		db = db.Where("creator = ?", data.Creator)
	}
	if data.Remarks != "" {
		db = db.Where("remarks like ?", "%"+data.Remarks+"%")
	}
	if data.Classify != 0 {
		db = db.Where("classify = ?", data.Classify)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询流程分页列表失败")
	return &list, total
}

func (m *workinfoModelImpl) FindList(data entity.FlowWorkInfo) *[]entity.FlowWorkInfo {
	list := make([]entity.FlowWorkInfo, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	db.Where("delete_time IS NULL")
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Creator != 0 {
		db = db.Where("creator = ?", data.Creator)
	}
	if data.Remarks != "" {
		db = db.Where("remarks like ?", "%"+data.Remarks+"%")
	}
	if data.Classify != 0 {
		db = db.Where("classify = ?", data.Classify)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询流程列表失败")
	return &list
}

func (m *workinfoModelImpl) Update(data entity.FlowWorkInfo) *entity.FlowWorkInfo {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改流程失败")
	return &data
}

func (m *workinfoModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.FlowWorkInfo{}, "id in (?)", ids).Error, "删除流程失败")
}
