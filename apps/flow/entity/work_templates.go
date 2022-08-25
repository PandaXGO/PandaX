package entity

import (
	"encoding/json"
	"github.com/XM-GO/PandaKit/model"
)

// FlowWorkTemplates 工作流表单模板
type FlowWorkTemplates struct {
	model.BaseAutoModel
	Name          string          `gorm:"column:name; type: varchar(128)" json:"name" binding:"required"`             // 模板名称
	FormStructure json.RawMessage `gorm:"column:form_structure; type: json" json:"form_structure" binding:"required"` // 表单结构
	Creator       int             `gorm:"column:creator; type: int(11)" json:"creator"`                               // 创建者
	Remarks       string          `gorm:"column:remarks; type: longtext" json:"remarks"`                              // 备注
}

func (FlowWorkTemplates) TableName() string {
	return "flow_work_templates"
}
