package entity

import "github.com/XM-GO/PandaKit/model"

type SysJob struct {
	JobId          int64  `json:"jobId" gorm:"primaryKey;autoIncrement"`    // 编码
	JobName        string `json:"jobName" gorm:"type:varchar(255);"`        // 名称
	JobGroup       string `json:"jobGroup" gorm:"type:varchar(255);"`       // 任务分组
	JobType        string `json:"jobType" gorm:"type:varchar(1);"`          // 任务类型
	CronExpression string `json:"cronExpression" gorm:"type:varchar(255);"` // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"type:varchar(255);"`   // 调用目标
	Args           string `json:"args" gorm:"type:varchar(255);"`           // 目标参数
	MisfirePolicy  string `json:"misfirePolicy" gorm:"type:varchar(1);"`    // 执行策略
	Concurrent     string `json:"concurrent" gorm:"type:varchar(1);"`       // 是否并发
	Status         string `json:"status" gorm:"type:varchar(1);"`           // 状态
	EntryId        int    `json:"entryId" gorm:"type:int;"`                 // job启动时返回的id
	CreateBy       string `json:"createBy" gorm:"type:varchar(128);comment:创建人"`
	UpdateBy       string `json:"updateBy" gorm:"type:varchar(128);comment:更新者"`
	model.BaseModel
}
