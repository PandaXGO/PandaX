package services

import (
	"pandax/base/biz"
	"pandax/base/global"
	"pandax/system/entity"
)

type (
	SysSettingsModel interface {
		Insert(data entity.SysSettings) *entity.SysSettings
		FindOne(key string) *entity.SysSettings
		FindList(data entity.SysSettings) *[]entity.SysSettings
		Update(data entity.SysSettings) *entity.SysSettings
		Delete(id []int64)
	}

	sysSettingsImpl struct {
		table string
	}
)

var SysSettingsModelDao SysSettingsModel = &sysSettingsImpl{
	table: `sys_settings`,
}

func (m *sysSettingsImpl) Insert(data entity.SysSettings) *entity.SysSettings {
	biz.ErrIsNil(global.Db.Table(m.table).Create(&data).Error, "添加配置失败")
	return &data
}

func (m *sysSettingsImpl) FindOne(key string) *entity.SysSettings {
	resData := new(entity.SysSettings)
	biz.ErrIsNil(global.Db.Table(m.table).Where("`key` = ?", key).First(resData).Error, "查询配置失败")
	return resData
}

func (m *sysSettingsImpl) FindList(data entity.SysSettings) *[]entity.SysSettings {
	list := make([]entity.SysSettings, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断

	biz.ErrIsNil(db.Find(&list).Error, "查询配置列表信息失败")
	return &list
}

func (m *sysSettingsImpl) Update(data entity.SysSettings) *entity.SysSettings {
	biz.ErrIsNil(global.Db.Table(m.table).Save(&data).Error, "修改配置信息")
	return &data
}

func (m *sysSettingsImpl) Delete(ids []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.SysSettings{}, "`id` in (?)", ids).Error, "删除配置信息失败")
	return
}
