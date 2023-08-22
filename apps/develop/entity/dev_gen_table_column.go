package entity

type DevGenTableColumn struct {
	ColumnId         int64  `gorm:"primaryKey;autoIncrement" json:"columnId"` // 编号
	TableId          int64  `gorm:"table_id"          json:"tableId"`         // 归属表编号
	TableName        string `gorm:"table_name" 		  json:"tableName"`
	ColumnName       string `gorm:"column_name"       json:"columnName"`    // 列名称
	ColumnComment    string `gorm:"column_comment"    json:"columnComment"` // 列描述
	ColumnType       string `gorm:"column_type"       json:"columnType"`    // 列类型
	ColumnKey        string `gorm:"column_key"        json:"columnKey"`
	GoType           string `gorm:"go_type"           json:"goType"`  // Go类型
	GoField          string `gorm:"go_field"          json:"goField"` // Go字段名
	JsonField        string `gorm:"json_field;" json:"jsonField"`
	HtmlField        string `gorm:"html_field"        json:"htmlField"`         // html字段名
	IsPk             string `gorm:"is_pk"             json:"isPk"`              // 是否主键（1是）
	IsIncrement      string `gorm:"is_increment"      json:"isIncrement"`       // 是否自增（1是）
	IsRequired       string `gorm:"is_required"       json:"isRequired"`        // 是否必填（1是）
	IsInsert         string `gorm:"is_insert"         json:"isInsert"`          // 是否为插入字段（1是）
	IsEdit           string `gorm:"is_edit"           json:"isEdit"`            // 是否编辑字段（1是）
	IsList           string `gorm:"is_list"           json:"isList"`            // 是否列表字段（1是）
	IsQuery          string `gorm:"is_query"          json:"isQuery"`           // 是否查询字段（1是）
	QueryType        string `gorm:"query_type"        json:"queryType"`         // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType         string `gorm:"html_type"         json:"htmlType"`          // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType         string `gorm:"dict_type"         json:"dictType"`          // 字典类型
	Sort             int    `gorm:"sort"              json:"sort"`              // 排序
	LinkTableName    string `gorm:"link_table_name"    json:"linkTableName"`    // 关联表名
	LinkTableClass   string `gorm:"link_table_class"   json:"linkTableClass"`   // 关联表类名
	LinkTablePackage string `gorm:"link_table_package" json:"linkTablePackage"` // 关联表包名
	LinkLabelId      string `gorm:"link_label_id"      json:"linkLabelId"`      // 关联表键名
	LinkLabelName    string `gorm:"link_label_name"    json:"linkLabelName"`    // 关联表字段值
}

type DBColumns struct {
	TableSchema            string `gorm:"column:TABLE_SCHEMA" json:"tableSchema"`
	TableName              string `gorm:"column:TABLE_NAME" json:"tableName"`
	ColumnName             string `gorm:"column:COLUMN_NAME" json:"columnName"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`
	IsNullable             string `gorm:"column:IS_NULLABLE" json:"isNullable"`
	DataType               string `gorm:"column:DATA_TYPE" json:"dataType"`
	CharacterMaximumLength string `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"characterMaximumLength"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME" json:"characterSetName"`
	ColumnType             string `gorm:"column:COLUMN_TYPE" json:"columnType"`
	ColumnKey              string `gorm:"column:COLUMN_KEY" json:"columnKey"`
	Extra                  string `gorm:"column:EXTRA" json:"extra"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT" json:"columnComment"`
}

type DBColumnsP struct {
	TableSchema            string `gorm:"column:table_schema" json:"tableSchema"`
	TableName              string `gorm:"column:table_name" json:"tableName"`
	ColumnName             string `gorm:"column:column_name" json:"columnName"`
	ColumnDefault          string `gorm:"column:column_default" json:"columnDefault"`
	IsNullable             string `gorm:"column:is_nullable" json:"isNullable"`
	DataType               string `gorm:"column:data_type" json:"dataType"`
	CharacterMaximumLength string `gorm:"column:character_maximum_length" json:"characterMaximumLength"`
	CharacterSetName       string `gorm:"column:character_set_name" json:"characterSetName"`
	ColumnType             string `gorm:"column:udt_name" json:"columnType"`
	ColumnKey              string ` json:"columnKey"` // 判断自增比较困难
	Extra                  string ` json:"extra"`
	ColumnComment          string ` json:"columnComment"`
}
