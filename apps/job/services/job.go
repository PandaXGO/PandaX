package services

import (
	"pandax/apps/job/entity"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
)

type (
	JobModel interface {
		Insert(data entity.SysJob) (*entity.SysJob, error)
		FindOne(jobId string) (*entity.SysJob, error)
		FindListPage(page, pageSize int, data entity.SysJob) (*[]entity.SysJob, int64, error)
		FindList(data entity.SysJob) (*[]entity.SysJob, error)
		Update(data entity.SysJob) (*entity.SysJob, error)
		Delete(jobId []string) error
		FindByEntryId(entryId int64) (*entity.SysJob, error)
		RemoveAllEntryID() error
		RemoveEntryID(EntryID int) error
	}

	jobModelImpl struct {
		table string
	}
)

var JobModelDao JobModel = &jobModelImpl{
	table: `jobs`,
}

func (m *jobModelImpl) Insert(data entity.SysJob) (*entity.SysJob, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *jobModelImpl) FindOne(jobId string) (*entity.SysJob, error) {
	resData := new(entity.SysJob)
	err := global.Db.Table(m.table).Where("id = ?", jobId).First(resData).Error
	return resData, err
}

func (m *jobModelImpl) FindListPage(page, pageSize int, data entity.SysJob) (*[]entity.SysJob, int64, error) {
	list := make([]entity.SysJob, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.JobName != "" {
		db = db.Where("job_name like ? ", "%"+data.JobName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}

	// 组织数据访问权限
	model.OrgAuthSet(db, data.RoleId, data.Owner)

	err := db.Count(&total).Error
	if err != nil {
		return &list, total, err
	}
	err = db.Order("create_time desc").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *jobModelImpl) FindList(data entity.SysJob) (*[]entity.SysJob, error) {
	list := make([]entity.SysJob, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.JobName != "" {
		db = db.Where("job_name like ? ", "%"+data.JobName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	// 组织数据访问权限
	err := model.OrgAuthSet(db, data.RoleId, data.Owner)
	if err != nil {
		return nil, err
	}
	err = db.Order("create_time desc").Find(&list).Error
	return &list, err
}

func (m *jobModelImpl) Update(data entity.SysJob) (*entity.SysJob, error) {
	err := global.Db.Table(m.table).Updates(&data).Error
	return &data, err
}

func (m *jobModelImpl) Delete(jobIds []string) error {
	err := global.Db.Table(m.table).Delete(&entity.SysJob{}, "id in (?)", jobIds).Error
	return err
}

func (m *jobModelImpl) FindByEntryId(entryId int64) (*entity.SysJob, error) {
	resData := new(entity.SysJob)
	err := global.Db.Table(m.table).Where("entry_id = ?", entryId).First(resData).Error
	return resData, err
}

func (m *jobModelImpl) RemoveAllEntryID() error {
	if err := global.Db.Table(m.table).Where("entry_id > ?", 0).Update("entry_id", 0).Error; err != nil {
		return err
	}
	return nil
}

func (m *jobModelImpl) RemoveEntryID(EntryID int) error {
	if err := global.Db.Table(m.table).Where("entry_id = ?", EntryID).Update("entry_id", 0).Error; err != nil {
		return err
	}
	return nil
}
