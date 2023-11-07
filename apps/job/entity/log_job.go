package entity

import (
	"pandax/pkg/global/model"
)

type JobLog struct {
	model.BaseAuthModel
	Name         string `json:"name" gorm:"type:varchar(128);comment:任务名称"`
	EntryId      int    `json:"entryId" gorm:"type:int;comment:任务id"`
	TargetInvoke string `json:"targetInvoke" gorm:"type:varchar(128);comment:调用方法"`
	LogInfo      string `json:"logInfo" gorm:"type:varchar(255);comment:日志信息"`
	Status       string `json:"status" gorm:"type:varchar(1);comment:状态"`

	RoleId int64 `gorm:"-"` // 角色数据权限
}
