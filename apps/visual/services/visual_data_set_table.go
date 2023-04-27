// ==========================================================================
// 生成日期：2023-04-10 03:05:21 +0800 CST
// 生成路径: apps/visual/services/visual_data_set_table.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualDataSetTableModel interface {
		Insert(data entity.VisualDataSetTable) *entity.VisualDataSetTable
		FindOne(tableId string) *entity.VisualDataSetTable
		FindListPage(page, pageSize int, data entity.VisualDataSetTable) (*[]entity.VisualDataSetTable, int64)
		FindList(data entity.VisualDataSetTable) *[]entity.VisualDataSetTable
		Update(data entity.VisualDataSetTable) *entity.VisualDataSetTable
		Delete(tableIds []string)
	}

	datasettableModelImpl struct {
		table string
	}
)

var VisualDataSetTableModelDao VisualDataSetTableModel = &datasettableModelImpl{
	table: `visual_data_set_table`,
}

func (m *datasettableModelImpl) Insert(data entity.VisualDataSetTable) *entity.VisualDataSetTable {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加数据集表失败")
	return &data
}

func (m *datasettableModelImpl) FindOne(tableId string) *entity.VisualDataSetTable {
	resData := new(entity.VisualDataSetTable)
	db := global.Db.Table(m.table).Where("table_id = ?", tableId)
	err := db.Preload("DataSource").First(resData).Error
	biz.ErrIsNil(err, "查询数据集表失败")
	return resData
}

func (m *datasettableModelImpl) FindListPage(page, pageSize int, data entity.VisualDataSetTable) (*[]entity.VisualDataSetTable, int64) {
	list := make([]entity.VisualDataSetTable, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DataSourceId != "" {
		db = db.Where("data_source_id = ?", data.DataSourceId)
	}
	if data.TableType != "" {
		db = db.Where("table_type = ?", data.TableType)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Preload("DataSource").Find(&list).Error
	biz.ErrIsNil(err, "查询数据集表分页列表失败")
	return &list, total
}

func (m *datasettableModelImpl) FindList(data entity.VisualDataSetTable) *[]entity.VisualDataSetTable {
	list := make([]entity.VisualDataSetTable, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DataSourceId != "" {
		db = db.Where("data_source_id = ?", data.DataSourceId)
	}
	if data.TableType != "" {
		db = db.Where("table_type = ?", data.TableType)
	}
	biz.ErrIsNil(db.Order("create_time").Preload("DataSource").Find(&list).Error, "查询数据集表列表失败")
	return &list
}

func (m *datasettableModelImpl) Update(data entity.VisualDataSetTable) *entity.VisualDataSetTable {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改数据集表失败")
	return &data
}

func (m *datasettableModelImpl) Delete(tableIds []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualDataSetTable{}, "table_id in (?)", tableIds).Error, "删除数据集表失败")
}
