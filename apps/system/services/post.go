package services

import (
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysPostModel interface {
		Insert(data entity.SysPost) (*entity.SysPost, error)
		FindOne(postId int64) (*entity.SysPost, error)
		FindListPage(page, pageSize int, data entity.SysPost) (*[]entity.SysPost, int64, error)
		FindList(data entity.SysPost) (*[]entity.SysPost, error)
		Update(data entity.SysPost) error
		Delete(postId []int64) error
	}

	sysPostModelImpl struct {
		table string
	}
)

var SysPostModelDao SysPostModel = &sysPostModelImpl{
	table: `sys_posts`,
}

func (m *sysPostModelImpl) Insert(data entity.SysPost) (*entity.SysPost, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *sysPostModelImpl) FindOne(postId int64) (*entity.SysPost, error) {
	resData := new(entity.SysPost)
	err := global.Db.Table(m.table).Where("post_id = ?", postId).First(resData).Error
	return resData, err
}

func (m *sysPostModelImpl) FindListPage(page, pageSize int, data entity.SysPost) (*[]entity.SysPost, int64, error) {
	list := make([]entity.SysPost, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.PostId != 0 {
		db = db.Where("post_id = ?", data.PostId)
	}
	if data.PostName != "" {
		db = db.Where("post_name like ?", "%"+data.PostName+"%")
	}
	if data.PostCode != "" {
		db = db.Where("post_code like ?", "%"+data.PostCode+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("sort").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *sysPostModelImpl) FindList(data entity.SysPost) (*[]entity.SysPost, error) {
	list := make([]entity.SysPost, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.PostId != 0 {
		db = db.Where("post_id = ?", data.PostId)
	}
	if data.PostName != "" {
		db = db.Where("post_name = ?", data.PostName)
	}
	if data.PostCode != "" {
		db = db.Where("post_code = ?", data.PostCode)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("sort").Find(&list).Error
	return &list, err
}

func (m *sysPostModelImpl) Update(data entity.SysPost) error {
	return global.Db.Table(m.table).Updates(&data).Error
}

func (m *sysPostModelImpl) Delete(postIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysPost{}, "post_id in (?)", postIds).Error
}
