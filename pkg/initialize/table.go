package initialize

import (
	"github.com/XM-GO/PandaKit/biz"
	devEntity "pandax/apps/develop/entity"
	flowEntity "pandax/apps/flow/entity"
	jobEntity "pandax/apps/job/entity"
	logEntity "pandax/apps/log/entity"
	resSourceEntity "pandax/apps/resource/entity"
	systemEntity "pandax/apps/system/entity"
	visualEntity "pandax/apps/visual/entity"
	"pandax/pkg/global"
)

// 初始化时如果没有表创建表
func InitTable() {
	m := global.Conf.Server
	if m.IsInitTable {
		biz.ErrIsNil(
			global.Db.AutoMigrate(
				//casbin.CasbinRule{},
				systemEntity.SysDept{},
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
				systemEntity.SysRoleDept{},
				systemEntity.SysNotice{},

				logEntity.LogLogin{},
				logEntity.LogOper{},
				logEntity.LogJob{},
				jobEntity.SysJob{},
				devEntity.DevGenTable{},
				devEntity.DevGenTableColumn{},
				resSourceEntity.ResOss{},
				resSourceEntity.ResEmail{},

				flowEntity.FlowWorkClassify{},
				flowEntity.FlowWorkInfo{},
				flowEntity.FlowWorkTemplates{},
				flowEntity.FlowWorkOrder{},
				flowEntity.FlowWorkOrderTemplate{},
				flowEntity.FlowWorkStage{},

				visualEntity.VisualDataSource{},
				visualEntity.VisualDataSetTable{},
				visualEntity.VisualDataSetField{},
				visualEntity.VisualScreen{},
				visualEntity.VisualScreenGroup{},
				visualEntity.VisualRuleChain{},
			),
			"初始化表失败")
	}
}
