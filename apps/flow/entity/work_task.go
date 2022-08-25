package entity

import (
	"github.com/XM-GO/PandaKit/model"
)

// FlowWorkTask 工作流任务
type FlowWorkTask struct {
	model.BaseAutoModel
	Name     string `gorm:"column:name; type: varchar(256)" json:"name"`          // 任务名称
	TaskType string `gorm:"column:task_type; type: varchar(45)" json:"task_type"` // 任务类型
	Content  string `gorm:"column:content; type: longtext" json:"content"`        // 任务内容
	Creator  int    `gorm:"column:creator; type: int(11)" json:"creator"`         // 创建者
	Remarks  string `gorm:"column:remarks; type: longtext" json:"remarks"`        // 备注
}

func (FlowWorkTask) TableName() string {
	return "flow_work_task"
}

// FlowWorkTaskHistory 工作流任务执行历史
type FlowWorkTaskHistory struct {
	model.BaseAutoModel
	Task          int    `gorm:"column:task; type: int(11)" json:"task"`                          // 任务ID
	Name          string `gorm:"column:name; type: varchar(256)" json:"name"`                     // 任务名称
	TaskType      int    `gorm:"column:task_type; type: int(11)" json:"task_type"`                // 任务类型, python, shell
	ExecutionTime string `gorm:"column:execution_time; type: varchar(128)" json:"execution_time"` // 执行时间
	Result        string `gorm:"column:result; type: longtext" json:"result"`                     // 任务返回
}

func (FlowWorkTaskHistory) TableName() string {
	return "flow_work_task_history"
}
