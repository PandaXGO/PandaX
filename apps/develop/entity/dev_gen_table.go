package entity

import "pandax/base/model"

type DevGenTable struct {
	TableId        int64               `gorm:"table_id,primary"   json:"tableId"`        // 编号
	TableName      string              `gorm:"table_name"         json:"tableName"`      // 表名称
	TableComment   string              `gorm:"table_comment"      json:"tableComment"`   // 表描述
	ClassName      string              `gorm:"class_name"         json:"className"`      // 实体类名称
	TplCategory    string              `gorm:"tpl_category"       json:"tplCategory"`    // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string              `gorm:"package_name"       json:"packageName"`    // 生成包路径
	ModuleName     string              `gorm:"module_name"        json:"moduleName"`     // 生成模块名
	BusinessName   string              `gorm:"business_name"      json:"businessName"`   // 生成业务名
	FunctionName   string              `gorm:"function_name"      json:"functionName"`   // 生成功能名
	FunctionAuthor string              `gorm:"function_author"    json:"functionAuthor"` // 生成功能作者
	Options        string              `gorm:"options"            json:"options"`        // 其它生成选项
	Remark         string              `gorm:"remark"             json:"remark"`         // 备注
	Columns        []DevGenTableColumn `gorm:"-" json:"columns"`                         // 字段信息
	PkColumn       DevGenTableColumn   `gorm:"-" json:"pkColumn"`                        // 主键列信息
	model.BaseModel
}

type ToolsGenTableExtend struct {
	DevGenTable
	TreeCode       string `gorm:"-" json:"treeCode"`       // 树编码字段
	TreeParentCode string `gorm:"-" json:"treeParentCode"` // 树父编码字段
	TreeName       string `gorm:"-" json:"treeName"`       // 树名称字段
}
