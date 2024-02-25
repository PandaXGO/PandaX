package services

import (
	"pandax/apps/job/entity"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
)

type (
	JobLogModel interface {
		Insert(data entity.JobLog) (*entity.JobLog, error)
		FindListPage(page, pageSize int, data entity.JobLog) (*[]entity.JobLog, int64, error)
		Delete(infoId []string) error
		DeleteAll() error
	}

	JobLogModelImpl struct {
		table string
	}
)

var JobLogModelDao JobLogModel = &JobLogModelImpl{
	table: `job_logs`,
}

func (m *JobLogModelImpl) Insert(data entity.JobLog) (*entity.JobLog, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *JobLogModelImpl) FindListPage(page, pageSize int, data entity.JobLog) (*[]entity.JobLog, int64, error) {
	list := make([]entity.JobLog, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}

	// 组织数据访问权限
	err := model.OrgAuthSet(db, data.RoleId, data.Owner)
	if err != nil {
		return &list, total, err
	}
	err = db.Count(&total).Error
	if err != nil {
		return &list, total, err
	}
	err = db.Order("create_time desc").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *JobLogModelImpl) Delete(logIds []string) error {
	err := global.Db.Table(m.table).Delete(&entity.JobLog{}, "id in (?)", logIds).Error
	return err
}

func (m *JobLogModelImpl) DeleteAll() error {
	return global.Db.Exec("DELETE FROM job_logs").Error
}
