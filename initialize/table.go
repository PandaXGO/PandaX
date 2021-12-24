package initialize

import (
	"pandax/apps/system/entity"

	jobEntity "pandax/apps/job/entity"
	logEntity "pandax/apps/log/entity"

	"pandax/base/biz"
	"pandax/base/config"
	"pandax/base/global"
)

// 初始化时如果没有表创建表
func InitTable() {
	m := config.Conf.Server
	if m.IsInitTable {
		biz.ErrIsNil(
			global.Db.AutoMigrate(
				entity.SysDept{},
				entity.SysApi{},
				entity.SysConfig{},
				entity.SysDictType{},
				entity.SysDictData{},
				logEntity.LogLogin{},
				logEntity.LogOper{},
				logEntity.LogJob{},
				entity.SysUser{},
				entity.SysRole{},
				entity.SysMenu{},
				entity.SysPost{},
				entity.SysRoleMenu{},
				entity.SysRoleDept{},
				entity.SysNotice{},
				jobEntity.SysJob{}),
			"初始化表失败")

	}
}
