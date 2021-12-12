package initialize

import (
	"pandax/base/biz"
	"pandax/base/casbin"
	"pandax/base/config"
	"pandax/base/global"
	"pandax/system/entity"
)

// 初始化时如果没有表创建表
func InitTable() {
	m := config.Conf.Server
	if m.IsInitTable {
		biz.ErrIsNil(
			global.Db.AutoMigrate(
				casbin.CasbinRule{},
				entity.SysDept{},
				entity.SysApi{},
				entity.SysConfig{},
				entity.SysDictType{},
				entity.SysDictData{},
				entity.LogLogin{},
				entity.LogOper{},
				entity.SysUser{},
				entity.SysRole{},
				entity.SysMenu{},
				entity.SysPost{},
				entity.SysRoleMenu{},
				entity.SysRoleDept{},
				entity.SysSettings{}), "初始化表失败")

	}
}
