package services

import (
	"github.com/PandaXGO/PandaKit/biz"
	"pandax/apps/job/entity"
	"pandax/pkg/global"
	"pandax/pkg/tool"
)

type (
	JobModel interface {
		Insert(data entity.SysJob) *entity.SysJob
		FindOne(jobId string) *entity.SysJob
		FindListPage(page, pageSize int, data entity.SysJob) (*[]entity.SysJob, int64)
		FindList(data entity.SysJob) *[]entity.SysJob
		Update(data entity.SysJob) *entity.SysJob
		Delete(jobId []string)
		FindByEntryId(entryId int64) *entity.SysJob
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

func (m *jobModelImpl) Insert(data entity.SysJob) *entity.SysJob {
	global.Db.Table(m.table).Create(&data)
	return &data
}

func (m *jobModelImpl) FindOne(jobId string) *entity.SysJob {
	resData := new(entity.SysJob)
	err := global.Db.Table(m.table).Where("id = ?", jobId).First(resData).Error
	biz.ErrIsNil(err, "查询任务信息失败")
	return resData
}

func (m *jobModelImpl) FindListPage(page, pageSize int, data entity.SysJob) (*[]entity.SysJob, int64) {
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
	tool.OrgAuthSet(db, data.RoleId, data.Owner)

	err := db.Count(&total).Error
	err = db.Order("create_time desc").Limit(pageSize).Offset(offset).Find(&list).Error

	biz.ErrIsNil(err, "查询任务分页信息失败")
	return &list, total
}

func (m *jobModelImpl) FindList(data entity.SysJob) *[]entity.SysJob {
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
	tool.OrgAuthSet(db, data.RoleId, data.Owner)
	err := db.Order("create_time desc").Find(&list).Error
	if err != nil {
		global.Log.Error("查询任务分页信息失败:" + err.Error())
	}
	return &list
}

func (m *jobModelImpl) Update(data entity.SysJob) *entity.SysJob {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改任务失败")
	return &data
}

func (m *jobModelImpl) Delete(jobIds []string) {
	err := global.Db.Table(m.table).Delete(&entity.SysJob{}, "id in (?)", jobIds).Error
	biz.ErrIsNil(err, "删除操作日志信息失败")
	return
}

func (m *jobModelImpl) FindByEntryId(entryId int64) *entity.SysJob {
	resData := new(entity.SysJob)
	err := global.Db.Table(m.table).Where("entry_id = ?", entryId).First(resData).Error
	biz.ErrIsNil(err, "查询失败")
	return resData
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
