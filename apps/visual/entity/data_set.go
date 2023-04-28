package entity

import "github.com/XM-GO/PandaKit/model"

type VisualDataSetTable struct {
	model.BaseModelD
	TableId      string           `gorm:"primary_key;tableId;comment:表id" json:"tableId"`
	DataSourceId string           `gorm:"dataSourceId;type:varchar(64);comment:数据源ID" json:"sourceId"`
	Name         string           `gorm:"name;type:varchar(64);comment:名称" json:"name"`
	TableType    string           `gorm:"tableType;type:varchar(64);comment:db,sql,excel,union" json:"tableType"`
	Info         string           `gorm:"info;type:TEXT;comment:原始表信息" json:"info"` //
	CreateBy     int64            `gorm:"create_by" json:"createBy"`                //创建人ID
	DataSource   VisualDataSource `gorm:"foreignKey:DataSourceId;references:SourceId" json:"dataSource"`
}

func (VisualDataSetTable) TableName() string {
	return "visual_data_set_table"
}

type VisualDataSetField struct {
	model.BaseModelD
	FieldId   string `gorm:"primary_key;fieldId;comment:表id" json:"fieldId"`
	TableId   string `gorm:"tableId;type:varchar(64);comment:表id" json:"tableId"`
	Name      string `gorm:"name;type:varchar(64);comment:字段描述" json:"name"`
	GroupType string `gorm:"group_type;type:varchar(1);comment:字段分组类型" json:"groupType"` //d 维度 g
	DeType    string `gorm:"de_type;type:varchar(1);comment:数据源字段类型：0-文本，1-时间，2-数值(包括小数)"   json:"deType"`
	DeName    string `gorm:"de_name;type:varchar(50);comment:数据源查询名"          json:"deName"`
	ExtField  string `gorm:"ext_field;type:varchar(1);comment:是否扩展字段 0否 1是" json:"extField"`
}

func (VisualDataSetField) TableName() string {
	return "visual_data_set_field"
}

type VisualDataSetRes struct {
	Data   []map[string]interface{} `json:"data"`
	Fields []string                 `json:"fields"`
}

type VisualDataSetRequest struct {
	SourceId string `json:"sourceId"`
	Info     string `json:"info"`
	Type     string `json:"type"`
}

type VisualDataSetTableFun struct {
	Name   string `gorm:"name;type:varchar(64);comment:函数名" json:"name"`
	DbType string `gorm:"db_type;type:varchar(50);comment:数据库类型" json:"dbType"` //d 维度 g
	Func   string `gorm:"func;type:varchar(64);comment:函数"   json:"func"`
	Desc   string `gorm:"desc;type:varchar(255);comment:描述"          json:"desc"`
}

func (VisualDataSetTableFun) TableName() string {
	return "visual_data_set_table_function"
}
