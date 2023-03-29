package entity

import "github.com/XM-GO/PandaKit/model"

// FlowWorkClassify 工作流流程分类
type FlowWorkClassify struct {
	model.BaseAutoModel
	Name    string `gorm:"column:name; type: varchar(128)" json:"name"` // 分类名称
	Creator int    `gorm:"column:creator; type: int" json:"creator"`    // 创建者
}

func (FlowWorkClassify) TableName() string {
	return "flow_work_classify"
}
