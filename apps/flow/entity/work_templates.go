package entity

import (
	"github.com/XM-GO/PandaKit/model"
)

// FlowWorkTemplates 工作流表单模板
type FlowWorkTemplates struct {
	model.BaseAutoModel
	Name          string      `gorm:"column:name; type: varchar(128)" json:"name" binding:"required"` // 模板名称
	FormStructure model.JSONB `gorm:"column:form_structure; type: json" json:"form_structure"`        // 表单结构
	Creator       int         `gorm:"column:creator; type: int" json:"creator"`                       // 创建者
	Remarks       string      `gorm:"column:remarks; type: text" json:"remarks"`                      // 备注
}

func (FlowWorkTemplates) TableName() string {
	return "flow_work_templates"
}
