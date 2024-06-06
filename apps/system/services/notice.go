package services

import (
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysNoticeModel interface {
		Insert(data entity.SysNotice) (*entity.SysNotice, error)
		FindOne(postId int64) (*entity.SysNotice, error)
		FindListPage(page, pageSize int, data entity.SysNotice) (*[]entity.SysNotice, int64, error)
		Update(data entity.SysNotice) error
		Delete(postId []int64) error
	}

	sysNoticeModelImpl struct {
		table string
	}
)

var SysNoticeModelDao SysNoticeModel = &sysNoticeModelImpl{
	table: `sys_notices`,
}

func (m *sysNoticeModelImpl) Insert(data entity.SysNotice) (*entity.SysNotice, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *sysNoticeModelImpl) FindOne(postId int64) (*entity.SysNotice, error) {
	resData := new(entity.SysNotice)
	err := global.Db.Table(m.table).Where("post_id = ?", postId).First(resData).Error
	return resData, err
}

func (m *sysNoticeModelImpl) FindListPage(page, pageSize int, data entity.SysNotice) (*[]entity.SysNotice, int64, error) {
	list := make([]entity.SysNotice, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Title != "" {
		db = db.Where("title like ?", "%"+data.Title+"%")
	}
	if data.NoticeType != "" {
		db = db.Where("notice_type = ?", data.NoticeType)
	}
	if len(data.OrganizationIds) > 0 {
		db = db.Where("organization_id in (?)", data.OrganizationIds)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *sysNoticeModelImpl) Update(data entity.SysNotice) error {
	return global.Db.Table(m.table).Updates(&data).Error
}

func (m *sysNoticeModelImpl) Delete(postIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysNotice{}, "notice_id in (?)", postIds).Error
}
