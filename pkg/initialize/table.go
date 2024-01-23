package initialize

import (
	devEntity "pandax/apps/develop/entity"
	jobEntity "pandax/apps/job/entity"
	logEntity "pandax/apps/log/entity"
	systemEntity "pandax/apps/system/entity"
	"pandax/kit/biz"
	"pandax/pkg/global"
)

// 初始化时如果没有表创建表
func InitTable() {
	m := global.Conf.Server
	if m.IsInitTable {
		biz.ErrIsNil(
			global.Db.AutoMigrate(
				//casbin.CasbinRule{},
				systemEntity.SysOrganization{},
				systemEntity.SysApi{},
				systemEntity.SysConfig{},
				systemEntity.SysDictType{},
				systemEntity.SysDictData{},
				systemEntity.SysUser{},
				systemEntity.SysTenants{},
				systemEntity.SysRole{},
				systemEntity.SysMenu{},
				systemEntity.SysPost{},
				systemEntity.SysRoleMenu{},
				systemEntity.SysRoleOrganization{},
				systemEntity.SysNotice{},

				logEntity.LogLogin{},
				logEntity.LogOper{},
				jobEntity.SysJob{},
				devEntity.DevGenTable{},
				devEntity.DevGenTableColumn{},
			),
			"初始化表失败")
	}
	err := global.TdDb.CreateEventTable()
	if err != nil {
		global.Log.Panic(err)
	}
}
