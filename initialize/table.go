package initialize

import (
	"pandax/apps/system/entity"

	logEntity "pandax/apps/log/entity"

	"pandax/base/biz"
	"pandax/base/casbin"
	"pandax/base/config"
	"pandax/base/global"
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
				logEntity.LogLogin{},
				logEntity.LogOper{},
				entity.SysUser{},
				entity.SysRole{},
				entity.SysMenu{},
				entity.SysPost{},
				entity.SysRoleMenu{},
				entity.SysRoleDept{}), "初始化表失败")

	}
}
