package services

import (
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysDictTypeModel interface {
		Insert(data entity.SysDictType) (*entity.SysDictType, error)
		FindOne(organizationId int64) (*entity.SysDictType, error)
		FindListPage(page, pageSize int, data entity.SysDictType) (*[]entity.SysDictType, int64, error)
		FindList(data entity.SysDictType) (*[]entity.SysDictType, error)
		Update(data entity.SysDictType) error
		Delete(organizationId []int64) error
	}

	sysDictTypeModelImpl struct {
		table string
	}
)

var SysDictTypeModelDao SysDictTypeModel = &sysDictTypeModelImpl{
	table: `sys_dict_types`,
}

func (m *sysDictTypeModelImpl) Insert(data entity.SysDictType) (*entity.SysDictType, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *sysDictTypeModelImpl) FindOne(dictId int64) (*entity.SysDictType, error) {
	resData := new(entity.SysDictType)
	err := global.Db.Table(m.table).Where("dict_id = ?", dictId).First(resData).Error
	return resData, err
}

func (m *sysDictTypeModelImpl) FindListPage(page, pageSize int, data entity.SysDictType) (*[]entity.SysDictType, int64, error) {
	list := make([]entity.SysDictType, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DictName != "" {
		db = db.Where("dict_name like ?", "%"+data.DictName+"%")
	}
	if data.DictType != "" {
		db = db.Where("dict_type like ?", "%"+data.DictType+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *sysDictTypeModelImpl) FindList(data entity.SysDictType) (*[]entity.SysDictType, error) {
	list := make([]entity.SysDictType, 0)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DictName != "" {
		db = db.Where("dict_name like ?", "%"+data.DictName+"%")
	}
	if data.DictType != "" {
		db = db.Where("dict_type like ?", "%"+data.DictType+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("create_time").Find(&list).Error
	return &list, err
}

func (m *sysDictTypeModelImpl) Update(data entity.SysDictType) error {
	return global.Db.Table(m.table).Where("dict_id = ?", data.DictId).Updates(&data).Error
}

func (m *sysDictTypeModelImpl) Delete(dictIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysDictType{}, "dict_id in (?)", dictIds).Error
}
