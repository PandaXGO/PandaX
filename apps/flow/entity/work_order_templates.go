package entity

import (
	"encoding/json"
	"github.com/XM-GO/PandaKit/model"
)

// FlowWorkOrderTemplate 工单绑定模版数据
type FlowWorkOrderTemplate struct {
	model.BaseAutoModel
	WorkOrder     int             `gorm:"column:work_order; type: int" json:"work_order"`          // 工单ID
	FormStructure json.RawMessage `gorm:"column:form_structure; type: json" json:"form_structure"` // 表单结构
	FormData      json.RawMessage `gorm:"column:form_data; type: json" json:"form_data"`           // 表单数据
}

func (FlowWorkOrderTemplate) TableName() string {
	return "flow_work_order_templates"
}
