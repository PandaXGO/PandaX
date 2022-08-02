package entity

import (
	"github.com/XM-GO/PandaKit/model"
)

type LogJob struct {
	LogId        int64  `json:"logId" gorm:"primary_key;AUTO_INCREMENT"` //主键
	Name         string `json:"name" gorm:"type:varchar(128);comment:任务名称"`
	JobGroup     string `json:"jobGroup" gorm:"type:varchar(128);comment:分组"`
	EntryId      int    `json:"entryId" gorm:"type:int;comment:任务id"`
	InvokeTarget string `json:"invokeTarget" gorm:"type:varchar(128);comment:调用方法"`
	LogInfo      string `json:"logInfo" gorm:"type:varchar(255);comment:日志信息"`
	Status       string `json:"status" gorm:"type:varchar(1);comment:状态"`
	model.BaseModel
}
