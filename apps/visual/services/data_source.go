// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/services/visual_data_source.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualDataSourceModel interface {
		Insert(data entity.VisualDataSource) *entity.VisualDataSource
		FindOne(sourceId string) *entity.VisualDataSource
		FindListPage(page, pageSize int, data entity.VisualDataSource) (*[]entity.VisualDataSource, int64)
		FindList(data entity.VisualDataSource) *[]entity.VisualDataSource
		Update(data entity.VisualDataSource) *entity.VisualDataSource
		Delete(sourceIds []string)
	}

	datasourceModelImpl struct {
		table string
	}
)

var VisualDataSourceModelDao VisualDataSourceModel = &datasourceModelImpl{
	table: `visual_data_source`,
}

func (m *datasourceModelImpl) Insert(data entity.VisualDataSource) *entity.VisualDataSource {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加数据源失败")
	return &data
}

func (m *datasourceModelImpl) FindOne(sourceId string) *entity.VisualDataSource {
	resData := new(entity.VisualDataSource)
	db := global.Db.Table(m.table).Where("source_id = ?", sourceId)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询数据源失败")
	return resData
}

func (m *datasourceModelImpl) FindListPage(page, pageSize int, data entity.VisualDataSource) (*[]entity.VisualDataSource, int64) {
	list := make([]entity.VisualDataSource, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.SourceId != "" {
		db = db.Where("source_id = ?", data.SourceId)
	}
	db.Where("delete_time IS NULL")
	if data.SourceComment != "" {
		db = db.Where("source_comment = ?", data.SourceComment)
	}
	if data.SourceType != "" {
		db = db.Where("source_type = ?", data.SourceType)
	}
	if data.SourceName != "" {
		db = db.Where("source_name like ?", "%"+data.SourceName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Configuration != "" {
		db = db.Where("configuration = ?", data.Configuration)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询数据源分页列表失败")
	return &list, total
}

func (m *datasourceModelImpl) FindList(data entity.VisualDataSource) *[]entity.VisualDataSource {
	list := make([]entity.VisualDataSource, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.SourceId != "" {
		db = db.Where("source_id = ?", data.SourceId)
	}
	db.Where("delete_time IS NULL")
	if data.SourceComment != "" {
		db = db.Where("source_comment = ?", data.SourceComment)
	}
	if data.SourceType != "" {
		db = db.Where("source_type = ?", data.SourceType)
	}
	if data.SourceName != "" {
		db = db.Where("source_name like ?", "%"+data.SourceName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Configuration != "" {
		db = db.Where("configuration = ?", data.Configuration)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询数据源列表失败")
	return &list
}

func (m *datasourceModelImpl) Update(data entity.VisualDataSource) *entity.VisualDataSource {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改数据源失败")
	return &data
}

func (m *datasourceModelImpl) Delete(sourceIds []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualDataSource{}, "source_id in (?)", sourceIds).Error, "删除数据源失败")
}
