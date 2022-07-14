package initialize

import (
	"pandax/apps/system/entity"
	//"pandax/base/casbin"

	devEntity "pandax/apps/develop/entity"
	jobEntity "pandax/apps/job/entity"
	logEntity "pandax/apps/log/entity"
	resSourceEntity "pandax/apps/resource/entity"

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
				//casbin.CasbinRule{},
				entity.SysDept{},
				entity.SysApi{},
				entity.SysConfig{},
				entity.SysDictType{},
				entity.SysDictData{},
				logEntity.LogLogin{},
				logEntity.LogOper{},
				logEntity.LogJob{},
				entity.SysUser{},
				entity.SysTenants{},
				entity.SysRole{},
				entity.SysMenu{},
				entity.SysPost{},
				entity.SysRoleMenu{},
				entity.SysRoleDept{},
				entity.SysNotice{},
				jobEntity.SysJob{},
				devEntity.DevGenTable{},
				devEntity.DevGenTableColumn{},
				resSourceEntity.ResOss{},
				resSourceEntity.ResEmail{},
			),
			"初始化表失败")

	}
}
