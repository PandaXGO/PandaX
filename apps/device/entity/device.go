package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"pandax/pkg/global/model"
	"time"
)

// DeviceGroup 设备分组
type DeviceGroup struct {
	model.BaseAuthModel
	Name        string `json:"name"  gorm:"type:varchar(128);comment:设备分组名称" validate:"required"`
	Pid         string `json:"pid" gorm:"type:varchar(64);comment:设备分组类型"`
	Path        string `json:"path" gorm:"type:varchar(255);comment:设备分组路径"`
	Description string `json:"description"  gorm:"type:varchar(255);comment:设备分组说明"`
	Sort        int64  `json:"sort" gorm:"type:int;comment:排序"`
	Status      string `gorm:"type:varchar(1);comment:状态" json:"status"`
	Ext         Ext    `json:"ext" gorm:"type:json;comment:扩展"` //可扩展的kv map,承载设备组的外围信息

	Children []DeviceGroup `json:"children" gorm:"-"` //子节点
}

type DeviceGroupLabel struct {
	Id       string             `gorm:"-" json:"id"`
	Name     string             `gorm:"-" json:"name"`
	Children []DeviceGroupLabel `gorm:"-" json:"children"`
}

type Device struct {
	model.BaseAuthModel
	Name        string    `json:"name"  gorm:"type:varchar(128);comment:设备名称" validate:"required,alphanum"` // mqtt 用户名英文
	ParentId    string    `json:"parentId" gorm:"type:varchar(64);comment:父设备"`
	DeviceType  string    `json:"deviceType" gorm:"type:varchar(64);comment:设备类型"`
	Token       string    `json:"token" gorm:"type:varchar(128);comment:设备token"`
	Alias       string    `json:"alias"  gorm:"type:varchar(128);comment:设备别名" `
	Pid         string    `json:"pid" gorm:"comment:产品Id" validate:"required"`
	Gid         string    `json:"gid" gorm:"comment:分组Id" validate:"required"`
	Description string    `json:"description"  gorm:"type:varchar(255);comment:说明"`
	Status      string    `json:"status" gorm:"type:varchar(1);comment:状态"`        //0 正常 1禁用
	LinkStatus  string    `json:"linkStatus" gorm:"type:varchar(10);comment:连接状态"` //inactive未激活 online在线 offline离线
	LastAt      time.Time `gorm:"column:last_time;comment:最后一次在线时间" json:"lastTime"`
	OtaVersion  string    `json:"otaVersion" gorm:"type:varchar(64);comment:固件版本" ` //上一次固件升级的版本
	Ext         Ext       `json:"ext" gorm:"type:json;comment:扩展"`                  //可扩展的kv map,承载设备组的外围信息

	Protocol string `json:"protocol" gorm:"-"`
}

type DeviceRes struct {
	Device
	DeviceGroup DeviceGroup `json:"deviceGroup" gorm:"foreignKey:Gid;references:Id"`
	Product     Product     `json:"product" gorm:"foreignKey:Pid;references:Id"`
}

type Ext map[string]interface{}

func (a Ext) Value() (driver.Value, error) {
	return json.Marshal(a)
}
func (a *Ext) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
