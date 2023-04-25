package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/XM-GO/PandaKit/model"
)

const (
	DbTypeMysql    = "MySQL"
	DbTypePostgres = "PostgreSQL"
)

type VisualDataSource struct {
	model.BaseModel
	SourceId      string   `gorm:"primary_key;source_id;comment:数据源Id"   json:"sourceId"`                          // 数据源Id
	SourceType    string   `gorm:"source_type;type:varchar(50);comment:数据源类型"         json:"sourceType"`           // 数据源类型
	SourceName    string   `gorm:"source_name;type:varchar(50);comment:数据源名称"      json:"sourceName"`              // 原名称
	SourceComment string   `gorm:"source_comment;type:varchar(50);comment:数据源描述"             json:"sourceComment"` // 描述
	Status        string   `gorm:"status;type:varchar(1);comment:数据源状态" json:"status"`
	Db            VisualDb `gorm:"db;type:text;comment:详细信息" json:"db"`
	CreateBy      int64    `gorm:"api" json:"createBy"` //创建人ID
}

type VisualDb struct {
	Host     string `gorm:"host" json:"host"`
	Port     int64  `gorm:"port" json:"port"`
	Dbname   string `gorm:"dbname" json:"dbname"`
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
	Config   string `gorm:"config" json:"config"` //额外的链接参数
	Schema   string `gorm:"schema" json:"schema"`
}

func (VisualDataSource) TableName() string {
	return "visual_data_source"
}

func (a VisualDb) Value() (driver.Value, error) {
	return json.Marshal(a)
}
func (a *VisualDb) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
