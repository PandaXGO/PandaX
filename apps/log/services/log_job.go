package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/log/entity"
	"pandax/pkg/global"
)

type (
	LogJobModel interface {
		Insert(data entity.LogJob) *entity.LogJob
		FindListPage(page, pageSize int, data entity.LogJob) (*[]entity.LogJob, int64)
		Delete(infoId []int64)
		DeleteAll()
	}

	logJobModelImpl struct {
		table string
	}
)

var LogJobModelDao LogJobModel = &logJobModelImpl{
	table: `log_jobs`,
}

func (m *logJobModelImpl) Insert(data entity.LogJob) *entity.LogJob {
	global.Db.Table(m.table).Create(&data)
	return &data
}

func (m *logJobModelImpl) FindListPage(page, pageSize int, data entity.LogJob) (*[]entity.LogJob, int64) {
	list := make([]entity.LogJob, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.JobGroup != "" {
		db = db.Where("job_group = ?", data.JobGroup)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	err := db.Where("delete_time IS NULL").Count(&total).Error
	err = db.Order("log_id desc").Limit(pageSize).Offset(offset).Find(&list).Error

	biz.ErrIsNil(err, "查询登录分页日志信息失败")
	return &list, total
}

func (m *logJobModelImpl) Delete(logIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.LogJob{}, "log_id in (?)", logIds).Error
	biz.ErrIsNil(err, "删除登录日志信息失败")
	return
}

func (m *logJobModelImpl) DeleteAll() {
	global.Db.Exec("DELETE FROM log_jobs")
}
