package entity

import (
	"encoding/json"
	"github.com/XM-GO/PandaKit/model"
)

// FlowWorkInfo 工作流信息
type FlowWorkInfo struct {
	model.BaseAutoModel
	Name        string          `gorm:"column:name; type:varchar(128)" json:"name"`                  // 流程名称
	Icon        string          `gorm:"column:icon; type:varchar(128)" json:"icon" `                 // 流程标签
	Structure   json.RawMessage `gorm:"column:structure; type:json" json:"structure" `               // 流程结构
	Classify    int             `gorm:"column:classify; type:int" json:"classify"`                   // 分类ID
	SubmitCount int             `gorm:"column:submit_count; type:int; default:0" json:"submitCount"` // 提交统计
	Creator     int             `gorm:"column:creator; type:int" json:"creator"`                     // 创建者
	Remarks     string          `gorm:"column:remarks; type:varchar(1024)" json:"remarks"`           // 流程备注
}

func (FlowWorkInfo) TableName() string {
	return "flow_work_info"
}
