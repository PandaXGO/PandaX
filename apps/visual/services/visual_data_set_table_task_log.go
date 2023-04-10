// ==========================================================================
// 生成日期：2023-04-10 11:31:34 +0800 CST
// 生成路径: apps/visual/services/visual_data_set_table_task_log.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualDataSetTableTaskLogModel interface {
		Insert(data entity.VisualDataSetTableTaskLog) *entity.VisualDataSetTableTaskLog
		FindOne(id string) *entity.VisualDataSetTableTaskLog
		FindListPage(page, pageSize int, data entity.VisualDataSetTableTaskLog) (*[]entity.VisualDataSetTableTaskLog, int64)
		FindList(data entity.VisualDataSetTableTaskLog) *[]entity.VisualDataSetTableTaskLog
		Update(data entity.VisualDataSetTableTaskLog) *entity.VisualDataSetTableTaskLog
		Delete(ids []string)
	}

	datasettabletasklogModelImpl struct {
		table string
	}
)

var VisualDataSetTableTaskLogModelDao VisualDataSetTableTaskLogModel = &datasettabletasklogModelImpl{
	table: `visual_data_set_table_task_log`,
}

func (m *datasettabletasklogModelImpl) Insert(data entity.VisualDataSetTableTaskLog) *entity.VisualDataSetTableTaskLog {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加生成数据集任务日志失败")
	return &data
}

func (m *datasettabletasklogModelImpl) FindOne(id string) *entity.VisualDataSetTableTaskLog {
	resData := new(entity.VisualDataSetTableTaskLog)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询生成数据集任务日志失败")
	return resData
}

func (m *datasettabletasklogModelImpl) FindListPage(page, pageSize int, data entity.VisualDataSetTableTaskLog) (*[]entity.VisualDataSetTableTaskLog, int64) {
	list := make([]entity.VisualDataSetTableTaskLog, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Id != "" {
		db = db.Where("id = ?", data.Id)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.TableId != "" {
		db = db.Where("table_id = ?", data.TableId)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询生成数据集任务日志分页列表失败")
	return &list, total
}

func (m *datasettabletasklogModelImpl) FindList(data entity.VisualDataSetTableTaskLog) *[]entity.VisualDataSetTableTaskLog {
	list := make([]entity.VisualDataSetTableTaskLog, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.TaskId != "" {
		db = db.Where("task_id = ?", data.TaskId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.TableId != "" {
		db = db.Where("table_id = ?", data.TableId)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询生成数据集任务日志列表失败")
	return &list
}

func (m *datasettabletasklogModelImpl) Update(data entity.VisualDataSetTableTaskLog) *entity.VisualDataSetTableTaskLog {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改生成数据集任务日志失败")
	return &data
}

func (m *datasettabletasklogModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualDataSetTableTaskLog{}, "id in (?)", ids).Error, "删除生成数据集任务日志失败")
}
