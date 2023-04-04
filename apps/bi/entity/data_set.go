package entity

import "github.com/XM-GO/PandaKit/model"

type DataSetGroup struct {
	model.BaseModelD
	Name  string `gorm:"name;type:varchar(64);comment:数据源类型" json:"name"`
	Pid   string `json:"pid"`
	Level int64  `json:"level"`
}

/*type DataSetTable struct {
	model.BaseModelD
	TableId      string              `gorm:"name;type:TEXT;comment:表id" json:"tableId"`
	DataSourceId string              `gorm:"name;type:TEXT;comment:数据圆ID" json:"data_source_Id"`
	TableType    string              `gorm:"name;type:TEXT;comment:db,sql,excel,custom,api" json:"tableType"` //'db,sql,excel,custom',
	Mode         string              `gorm:"name;type:TEXT;comment:原始表信息" json:"mode"`                        //'连接模式：0-直连，1-定时同步',
	Info         string              `gorm:"name;type:TEXT;comment:原始表信息" json:"name"`
	CreateBy     int64               `gorm:"create_by" json:"createBy"`         //创建人ID
	Columns      []DevGenTableColumn `gorm:"-"                  json:"columns"` //字段信息
}
*/
