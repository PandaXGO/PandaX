package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysDictDataModel interface {
		Insert(data entity.SysDictData) *entity.SysDictData
		FindOne(dictCode int64) *entity.SysDictData
		FindListPage(page, pageSize int, data entity.SysDictData) (*[]entity.SysDictData, int64)
		FindList(data entity.SysDictData) *[]entity.SysDictData
		Update(data entity.SysDictData) *entity.SysDictData
		Delete(dictCode []int64)
	}

	sysDictDataModelImpl struct {
		table string
	}
)

var SysDictDataModelDao SysDictDataModel = &sysDictDataModelImpl{
	table: `sys_dict_data`,
}

func (m *sysDictDataModelImpl) Insert(data entity.SysDictData) *entity.SysDictData {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "新增字典数据失败")
	return &data
}

func (m *sysDictDataModelImpl) FindOne(codeId int64) *entity.SysDictData {
	resData := new(entity.SysDictData)
	err := global.Db.Table(m.table).Where("dict_code = ?", codeId).First(resData).Error
	biz.ErrIsNil(err, "查询字典数据信息失败")
	return resData
}

func (m *sysDictDataModelImpl) FindListPage(page, pageSize int, data entity.SysDictData) (*[]entity.SysDictData, int64) {
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
	err = db.Order("dict_sort").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询字典数据分页列表信息失败")
	return &list, total
}

func (m *sysDictDataModelImpl) FindList(data entity.SysDictData) *[]entity.SysDictData {
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
	biz.ErrIsNil(err, "查询字典数据列表信息失败")
	return &list
}

func (m *sysDictDataModelImpl) Update(data entity.SysDictData) *entity.SysDictData {
	err := global.Db.Table(m.table).Where("dict_code = ?", data.DictCode).Updates(&data).Error
	biz.ErrIsNil(err, "修改字典数据信息失败")
	return &data
}

func (m *sysDictDataModelImpl) Delete(codeIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.SysDept{}, "dict_code in (?)", codeIds).Error
	biz.ErrIsNil(err, "删除字典数据信息失败")
	return
}
