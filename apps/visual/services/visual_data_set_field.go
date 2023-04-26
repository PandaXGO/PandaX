// ==========================================================================
// 生成日期：2023-04-10 03:05:21 +0800 CST
// 生成路径: apps/visual/services/visual_data_set_field.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualDataSetFieldModel interface {
		Insert(data entity.VisualDataSetField) *entity.VisualDataSetField
		FindOne(fieldId string) *entity.VisualDataSetField
		FindListPage(page, pageSize int, data entity.VisualDataSetField) (*[]entity.VisualDataSetField, int64)
		FindList(data entity.VisualDataSetField) *[]entity.VisualDataSetField
		Update(data entity.VisualDataSetField) *entity.VisualDataSetField
		Delete(fieldIds []string)
	}

	datasetfieldModelImpl struct {
		table string
	}
)

var VisualDataSetFieldModelDao VisualDataSetFieldModel = &datasetfieldModelImpl{
	table: `visual_data_set_field`,
}

func (m *datasetfieldModelImpl) Insert(data entity.VisualDataSetField) *entity.VisualDataSetField {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加数据集字段失败")
	return &data
}

func (m *datasetfieldModelImpl) FindOne(fieldId string) *entity.VisualDataSetField {
	resData := new(entity.VisualDataSetField)
	db := global.Db.Table(m.table).Where("field_id = ?", fieldId)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询数据集字段失败")
	return resData
}

func (m *datasetfieldModelImpl) FindListPage(page, pageSize int, data entity.VisualDataSetField) (*[]entity.VisualDataSetField, int64) {
	list := make([]entity.VisualDataSetField, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.TableId != "" {
		db = db.Where("table_id = ?", data.TableId)
	}
	if data.GroupType != "" {
		db = db.Where("group_type = ?", data.GroupType)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询数据集字段分页列表失败")
	return &list, total
}

func (m *datasetfieldModelImpl) FindList(data entity.VisualDataSetField) *[]entity.VisualDataSetField {
	list := make([]entity.VisualDataSetField, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.TableId != "" {
		db = db.Where("table_id = ?", data.TableId)
	}
	if data.GroupType != "" {
		db = db.Where("group_type = ?", data.GroupType)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询数据集字段列表失败")
	return &list
}

func (m *datasetfieldModelImpl) Update(data entity.VisualDataSetField) *entity.VisualDataSetField {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改数据集字段失败")
	return &data
}

func (m *datasetfieldModelImpl) Delete(fieldIds []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualDataSetField{}, "field_id in (?)", fieldIds).Error, "删除数据集字段失败")
}
