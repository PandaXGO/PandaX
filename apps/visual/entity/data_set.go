package entity

import "github.com/XM-GO/PandaKit/model"

// VisualDataSetGroup 数据集分组
type VisualDataSetGroup struct {
	model.BaseModelD
	Id    string `gorm:"primary_key;id;type:varchar(64);" json:"id"`
	Name  string `gorm:"name;type:varchar(64);comment:数据源类型" json:"name"`
	Pid   string `gorm:"pid;type:varchar(64);comment:父Id" json:"pid"`
	Level int64  `gorm:"level;type:int;comment:等级" json:"level"`
}

func (VisualDataSetGroup) TableName() string {
	return "visual_data_set_group"
}

type VisualDataSetTable struct {
	model.BaseModelD
	TableId      string               `gorm:"primary_key;tableId;comment:表id" json:"tableId"`
	DataSourceId string               `gorm:"dataSourceId;type:varchar(64);comment:数据圆ID" json:"dataSourceId"`
	TableType    string               `gorm:"tableType;type:varchar(64);comment:db,sql,excel,custom,api" json:"tableType"` //'db,sql,excel,custom',
	Mode         string               `gorm:"mode;type:varchar(1);comment:原始表信息" json:"mode"`                              //'连接模式：0-直连，1-定时同步',
	Info         string               `gorm:"info;type:TEXT;comment:原始表信息" json:"info"`
	CreateBy     int64                `gorm:"create_by" json:"createBy"`        //创建人ID
	Fields       []VisualDataSetField `gorm:"-"                  json:"fields"` //字段信息
}

func (VisualDataSetTable) TableName() string {
	return "visual_data_set_table"
}

type VisualDataSetField struct {
	model.BaseModelD
	FieldId   string `gorm:"primary_key;fieldId;comment:表id" json:"fieldId"`
	TableId   string `gorm:"tableId;type:varchar(64);comment:表id" json:"tableId"`
	Comment   string `gorm:"comment;type:varchar(64);comment:字段描述"    json:"columnComment"` // 列描述
	Name      string `gorm:"name;type:varchar(64);comment:字段名" json:"name"`
	Type      string `gorm:"type;type:varchar(50);comment:字段类型" json:"type"`
	GoType    string `gorm:"go_type;type:varchar(50);comment:go字段类型"           json:"goType"`  // Go类型
	GoField   string `gorm:"go_field;type:varchar(50);comment:go对应字段"          json:"goField"` // Go字段名
	JsonField string `gorm:"json_field;type:varchar(50);comment:json对应字段" json:"jsonField"`
}

func (VisualDataSetField) TableName() string {
	return "visual_data_set_field"
}
