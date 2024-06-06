package services

import (
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysConfigModel interface {
		Insert(data entity.SysConfig) (*entity.SysConfig, error)
		FindOne(dictCode int64) (*entity.SysConfig, error)
		FindListPage(page, pageSize int, data entity.SysConfig) (*[]entity.SysConfig, int64, error)
		FindList(data entity.SysConfig) (*[]entity.SysConfig, error)
		Update(data entity.SysConfig) error
		Delete(dictCode []int64) error
	}

	sysSysConfigModelImpl struct {
		table string
	}
)

var SysSysConfigModelDao SysConfigModel = &sysSysConfigModelImpl{
	table: `sys_configs`,
}

func (m *sysSysConfigModelImpl) Insert(data entity.SysConfig) (*entity.SysConfig, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *sysSysConfigModelImpl) FindOne(configId int64) (*entity.SysConfig, error) {
	resData := new(entity.SysConfig)
	err := global.Db.Table(m.table).Where("config_id = ?", configId).First(resData).Error
	return resData, err
}

func (m *sysSysConfigModelImpl) FindListPage(page, pageSize int, data entity.SysConfig) (*[]entity.SysConfig, int64, error) {
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
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *sysSysConfigModelImpl) FindList(data entity.SysConfig) (*[]entity.SysConfig, error) {
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
	return &list, err
}

func (m *sysSysConfigModelImpl) Update(data entity.SysConfig) error {
	return global.Db.Table(m.table).Model(&data).Updates(&data).Error
}

func (m *sysSysConfigModelImpl) Delete(configIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysConfig{}, "config_id in (?)", configIds).Error
}
