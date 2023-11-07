package entity

import (
	"pandax/pkg/global/model"
)

type SysJob struct {
	model.BaseAuthModel
	JobName        string `json:"jobName" gorm:"type:varchar(255);"`                 // 名称
	TargetInvoke   string `json:"targetInvoke" gorm:"type:varchar(64);comment:目标类型"` //调用目标 设备还是产品
	TargetArgs     string `json:"targetArgs" gorm:"type:varchar(64);comment:目标Id"`   //目标传参 设备或者产品id
	JobContent     string `json:"jobContent" gorm:"type:json;comment:任务内容"`          //目标传参 要执行的内容
	CronExpression string `json:"cronExpression" gorm:"type:varchar(255);"`          // cron表达式
	MisfirePolicy  string `json:"misfirePolicy" gorm:"type:varchar(1);"`             // 执行策略
	Status         string `json:"status" gorm:"type:varchar(1);"`                    // 状态
	EntryId        int    `json:"entryId" gorm:"type:int;"`                          // job启动时返回的id

	RoleId int64 `gorm:"-"` // 角色数据权限
}
