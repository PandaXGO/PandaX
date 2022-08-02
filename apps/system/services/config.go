package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysConfigModel interface {
		Insert(data entity.SysConfig) *entity.SysConfig
		FindOne(dictCode int64) *entity.SysConfig
		FindListPage(page, pageSize int, data entity.SysConfig) (*[]entity.SysConfig, int64)
		FindList(data entity.SysConfig) *[]entity.SysConfig
		Update(data entity.SysConfig) *entity.SysConfig
		Delete(dictCode []int64)
	}

	sysSysConfigModelImpl struct {
		table string
	}
)

var SysSysConfigModelDao SysConfigModel = &sysSysConfigModelImpl{
	table: `sys_configs`,
}

func (m *sysSysConfigModelImpl) Insert(data entity.SysConfig) *entity.SysConfig {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "新增配置失败")
	return &data
}

func (m *sysSysConfigModelImpl) FindOne(configId int64) *entity.SysConfig {
	resData := new(entity.SysConfig)
	err := global.Db.Table(m.table).Where("config_id = ?", configId).First(resData).Error
	biz.ErrIsNil(err, "查询配置信息失败")
	return resData
}

func (m *sysSysConfigModelImpl) FindListPage(page, pageSize int, data entity.SysConfig) (*[]entity.SysConfig, int64) {
	list := make([]entity.SysConfig, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.ConfigName != "" {
		db = db.Where("config_name like ?", "%"+data.ConfigName+"%")
	}
	if data.ConfigKey != "" {
		db = db.Where("config_key like ?", "%"+data.ConfigKey+"%")
	}
	if data.ConfigType != "" {
		db = db.Where("config_type = ?", data.ConfigType)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询配置分页列表信息失败")
	return &list, total
}

func (m *sysSysConfigModelImpl) FindList(data entity.SysConfig) *[]entity.SysConfig {
	list := make([]entity.SysConfig, 0)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.ConfigName != "" {
		db = db.Where("config_name like ?", "%"+data.ConfigName+"%")
	}
	if data.ConfigKey != "" {
		db = db.Where("config_key like ?", "%"+data.ConfigKey+"%")
	}
	if data.ConfigType != "" {
		db = db.Where("config_type = ?", data.ConfigType)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("create_time").Find(&list).Error
	biz.ErrIsNil(err, "查询配置列表信息失败")
	return &list
}

func (m *sysSysConfigModelImpl) Update(data entity.SysConfig) *entity.SysConfig {
	err := global.Db.Table(m.table).Model(&data).Updates(&data).Error
	biz.ErrIsNil(err, "修改配置信息失败")
	return &data
}

func (m *sysSysConfigModelImpl) Delete(configIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.SysConfig{}, "config_id in (?)", configIds).Error
	biz.ErrIsNil(err, "删除配置信息失败")
	return
}
