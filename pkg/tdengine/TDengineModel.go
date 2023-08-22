package tdengine

import "time"

type TDEngineTablesList struct {
	TableName  string     `json:"tableName"        description:"表名"`
	DbName     string     `json:"dbName"        description:"数据库名"`
	StableName string     `json:"stableName"        description:"超级表名"`
	CreateTime *time.Time `json:"createTime" description:"创建时间"`
}

type TDEngineTableInfo struct {
	Field  string `json:"field"        description:"字段名"`
	Type   string `json:"type"        description:"类型"`
	Length int    `json:"length"        description:"长度"`
	Note   string `json:"note" description:"note"`
}

type TableDataInfo struct {
	Filed []string                 `json:"filed"        description:"字段"`
	Info  []map[string]interface{} `json:"info"        description:"数据"`
}

// 日志 TDengine
type TdLog struct {
	Ts      string `json:"ts" dc:"时间"`
	Device  string `json:"device" dc:"设备标识"`
	Type    string `json:"type" dc:"日志类型"`
	Content string `json:"content" dc:"日志内容"`
}
