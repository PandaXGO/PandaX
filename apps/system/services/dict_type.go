package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysDictTypeModel interface {
		Insert(data entity.SysDictType) *entity.SysDictType
		FindOne(deptId int64) *entity.SysDictType
		FindListPage(page, pageSize int, data entity.SysDictType) (*[]entity.SysDictType, int64)
		FindList(data entity.SysDictType) *[]entity.SysDictType
		Update(data entity.SysDictType) *entity.SysDictType
		Delete(deptId []int64)
	}

	sysDictTypeModelImpl struct {
		table string
	}
)

var SysDictTypeModelDao SysDictTypeModel = &sysDictTypeModelImpl{
	table: `sys_dict_types`,
}

func (m *sysDictTypeModelImpl) Insert(data entity.SysDictType) *entity.SysDictType {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "新增字典类型失败")
	return &data
}

func (m *sysDictTypeModelImpl) FindOne(dictId int64) *entity.SysDictType {
	resData := new(entity.SysDictType)
	err := global.Db.Table(m.table).Where("dict_id = ?", dictId).First(resData).Error
	biz.ErrIsNil(err, "查询字典类型信息失败")
	return resData
}

func (m *sysDictTypeModelImpl) FindListPage(page, pageSize int, data entity.SysDictType) (*[]entity.SysDictType, int64) {
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
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询字典类型分页列表信息失败")
	return &list, total
}

func (m *sysDictTypeModelImpl) FindList(data entity.SysDictType) *[]entity.SysDictType {
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
	biz.ErrIsNil(err, "查询字典类型列表信息失败")
	return &list
}

func (m *sysDictTypeModelImpl) Update(data entity.SysDictType) *entity.SysDictType {
	err := global.Db.Table(m.table).Where("dict_id = ?", data.DictId).Updates(&data).Error
	biz.ErrIsNil(err, "修改字典类型信息失败")
	return &data
}

func (m *sysDictTypeModelImpl) Delete(dictIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.SysDictType{}, "dict_id in (?)", dictIds).Error
	biz.ErrIsNil(err, "删除字典类型信息失败")
	return
}
