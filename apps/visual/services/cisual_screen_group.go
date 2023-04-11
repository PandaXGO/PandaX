// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/services/visual_data_set_group.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualScreenGroupModel interface {
		Insert(data entity.VisualScreenGroup) *entity.VisualScreenGroup
		FindOne(id string) *entity.VisualScreenGroup
		FindListPage(page, pageSize int, data entity.VisualScreenGroup) (*[]entity.VisualScreenGroup, int64)
		FindList(data entity.VisualScreenGroup) *[]entity.VisualScreenGroup
		Update(data entity.VisualScreenGroup) *entity.VisualScreenGroup
		Delete(ids []string)
	}

	screenGroupModelImpl struct {
		table string
	}
)

var VisualScreenGroupModelDao VisualScreenGroupModel = &screenGroupModelImpl{
	table: `visual_screen_group`,
}

func (m *screenGroupModelImpl) Insert(data entity.VisualScreenGroup) *entity.VisualScreenGroup {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加数据集分组失败")
	return &data
}

func (m *screenGroupModelImpl) FindOne(id string) *entity.VisualScreenGroup {
	resData := new(entity.VisualScreenGroup)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询数据集分组失败")
	return resData
}

func (m *screenGroupModelImpl) FindListPage(page, pageSize int, data entity.VisualScreenGroup) (*[]entity.VisualScreenGroup, int64) {
	list := make([]entity.VisualScreenGroup, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询数据集分组分页列表失败")
	return &list, total
}

func (m *screenGroupModelImpl) FindList(data entity.VisualScreenGroup) *[]entity.VisualScreenGroup {
	list := make([]entity.VisualScreenGroup, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询数据集分组列表失败")
	return &list
}

func (m *screenGroupModelImpl) Update(data entity.VisualScreenGroup) *entity.VisualScreenGroup {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改数据集分组失败")
	return &data
}

func (m *screenGroupModelImpl) Delete(s []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualScreenGroup{}, "id in (?)", s).Error, "删除数据集分组失败")
}
