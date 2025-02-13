package services

import (
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysDictDataModel interface {
		Insert(data entity.SysDictData) (*entity.SysDictData, error)
		FindOne(dictCode int64) (*entity.SysDictData, error)
		FindListPage(page, pageSize int, data entity.SysDictData) (*[]entity.SysDictData, int64, error)
		FindList(data entity.SysDictData) (*[]entity.SysDictData, error)
		Update(data entity.SysDictData) error
		Delete(dictCode []int64) error
	}

	sysDictDataModelImpl struct {
		table string
	}
)

var SysDictDataModelDao SysDictDataModel = &sysDictDataModelImpl{
	table: `sys_dict_data`,
}

func (m *sysDictDataModelImpl) Insert(data entity.SysDictData) (*entity.SysDictData, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *sysDictDataModelImpl) FindOne(codeId int64) (*entity.SysDictData, error) {
	resData := new(entity.SysDictData)
	err := global.Db.Table(m.table).Where("dict_code = ?", codeId).First(resData).Error
	return resData, err
}

func (m *sysDictDataModelImpl) FindListPage(page, pageSize int, data entity.SysDictData) (*[]entity.SysDictData, int64, error) {
	list := make([]entity.SysDictData, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DictLabel != "" {
		db = db.Where("dict_label = ?", data.DictLabel)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.DictType != "" {
		db = db.Where("dict_type = ?", data.DictType)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("dict_sort").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *sysDictDataModelImpl) FindList(data entity.SysDictData) (*[]entity.SysDictData, error) {
	list := make([]entity.SysDictData, 0)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DictLabel != "" {
		db = db.Where("dict_label like ?", "%"+data.DictLabel+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.DictType != "" {
		db = db.Where("dict_type = ?", data.DictType)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("dict_sort").Find(&list).Error
	return &list, err
}

func (m *sysDictDataModelImpl) Update(data entity.SysDictData) error {
	return global.Db.Table(m.table).Where("dict_code = ?", data.DictCode).Updates(&data).Error
}

func (m *sysDictDataModelImpl) Delete(codeIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysDictData{}, "dict_code in (?)", codeIds).Error
}
