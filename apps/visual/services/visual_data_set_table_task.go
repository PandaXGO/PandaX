// ==========================================================================
// 生成日期：2023-04-10 11:31:34 +0800 CST
// 生成路径: apps/visual/services/visual_data_set_table_task.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualDataSetTableTaskModel interface {
		Insert(data entity.VisualDataSetTableTask) *entity.VisualDataSetTableTask
		FindOne(id string) *entity.VisualDataSetTableTask
		FindListPage(page, pageSize int, data entity.VisualDataSetTableTask) (*[]entity.VisualDataSetTableTask, int64)
		FindList(data entity.VisualDataSetTableTask) *[]entity.VisualDataSetTableTask
		Update(data entity.VisualDataSetTableTask) *entity.VisualDataSetTableTask
		Delete(ids []string)
	}

	datasettabletaskModelImpl struct {
		table string
	}
)

var VisualDataSetTableTaskModelDao VisualDataSetTableTaskModel = &datasettabletaskModelImpl{
	table: `visual_data_set_table_task`,
}

func (m *datasettabletaskModelImpl) Insert(data entity.VisualDataSetTableTask) *entity.VisualDataSetTableTask {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加生成数据集任务失败")
	return &data
}

func (m *datasettabletaskModelImpl) FindOne(id string) *entity.VisualDataSetTableTask {
	resData := new(entity.VisualDataSetTableTask)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询生成数据集任务失败")
	return resData
}

func (m *datasettabletaskModelImpl) FindListPage(page, pageSize int, data entity.VisualDataSetTableTask) (*[]entity.VisualDataSetTableTask, int64) {
	list := make([]entity.VisualDataSetTableTask, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.TableId != "" {
		db = db.Where("table_id = ?", data.TableId)
	}
	if data.Type != "" {
		db = db.Where("type = ?", data.Type)
	}
	if data.Rate != "" {
		db = db.Where("rate = ?", data.Rate)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询生成数据集任务分页列表失败")
	return &list, total
}

func (m *datasettabletaskModelImpl) FindList(data entity.VisualDataSetTableTask) *[]entity.VisualDataSetTableTask {
	list := make([]entity.VisualDataSetTableTask, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.TableId != "" {
		db = db.Where("table_id = ?", data.TableId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Type != "" {
		db = db.Where("type = ?", data.Type)
	}
	if data.Rate != "" {
		db = db.Where("rate = ?", data.Rate)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询生成数据集任务列表失败")
	return &list
}

func (m *datasettabletaskModelImpl) Update(data entity.VisualDataSetTableTask) *entity.VisualDataSetTableTask {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改生成数据集任务失败")
	return &data
}

func (m *datasettabletaskModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualDataSetTableTask{}, "id in (?)", ids).Error, "删除生成数据集任务失败")
}
