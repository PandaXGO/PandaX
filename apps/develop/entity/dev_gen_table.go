package entity

import "github.com/XM-GO/PandaKit/model"

type DevGenTable struct {
	TableId        int64               `gorm:"primaryKey;autoIncrement"   json:"tableId"` // 编号
	TableName      string              `gorm:"table_name"         json:"tableName"`       // 表名称
	TableComment   string              `gorm:"table_comment"      json:"tableComment"`    // 表描述
	ClassName      string              `gorm:"class_name"         json:"className"`       // 实体类名称
	TplCategory    string              `gorm:"tpl_category"       json:"tplCategory"`     // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string              `gorm:"package_name"       json:"packageName"`     // 生成包路径
	ModuleName     string              `gorm:"module_name"        json:"moduleName"`      // 生成模块名
	BusinessName   string              `gorm:"business_name"      json:"businessName"`    // 生成业务名
	FunctionName   string              `gorm:"function_name"      json:"functionName"`    // 生成功能名
	FunctionAuthor string              `gorm:"function_author"    json:"functionAuthor"`  // 生成功能作者
	Options        string              `gorm:"options"            json:"options"`         // 其它生成选项
	Remark         string              `gorm:"remark"             json:"remark"`          // 备注
	PkColumn       string              `gorm:"pk_column;"         json:"pkColumn"`
	PkGoField      string              `gorm:"pk_go_field"        json:"pkGoField"`
	PkJsonField    string              `gorm:"pk_json_field"      json:"pkJsonField"`
	Columns        []DevGenTableColumn `gorm:"-"                  json:"columns"` // 字段信息
	model.BaseModel
}

type DBTables struct {
	TableName      string `gorm:"column:TABLE_NAME" json:"tableName"`
	Engine         string `gorm:"column:ENGINE" json:"engine"`
	TableRows      string `gorm:"column:TABLE_ROWS" json:"tableRows"`
	TableCollation string `gorm:"column:TABLE_COLLATION" json:"tableCollation"`
	CreateTime     string `gorm:"column:CREATE_TIME" json:"createTime"`
	UpdateTime     string `gorm:"column:UPDATE_TIME" json:"updateTime"`
	TableComment   string `gorm:"column:TABLE_COMMENT" json:"tableComment"`
}
