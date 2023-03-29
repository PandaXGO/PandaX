package entity

import (
	"github.com/XM-GO/PandaKit/model"
)

// FlowWorkStage 工作流工序（流转历史）
type FlowWorkStage struct {
	model.BaseAutoModel
	Title        string `gorm:"column:title; type: varchar(128)" json:"title"`        // 工单标题
	WorkOrder    int    `gorm:"column:work_order; type: int" json:"work_order"`       // 工单ID
	State        string `gorm:"column:state; type: varchar(128)" json:"state"`        // 工单状态
	Source       string `gorm:"column:source; type: varchar(128)" json:"source"`      // 源节点ID
	Target       string `gorm:"column:target; type: varchar(128)" json:"target"`      // 目标节点ID
	Stage        string `gorm:"column:stage; type: varchar(128)" json:"stage"`        // 流转ID
	Status       int    `gorm:"column:status; type: int" json:"status"`               // 流转状态 1 同意， 0 拒绝， 2 其他
	Processor    string `gorm:"column:processor; type: varchar(45)" json:"processor"` // 处理人
	ProcessorId  int    `gorm:"column:processor_id; type: int" json:"processor_id"`   // 处理人ID
	CostDuration int64  `gorm:"column:cost_duration; type: int" json:"cost_duration"` // 处理时长
	Remarks      string `gorm:"column:remarks; type: text" json:"remarks"`            // 备注
}

func (FlowWorkStage) TableName() string {
	return "flow_work_stage"
}
