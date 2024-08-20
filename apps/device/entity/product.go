package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"pandax/pkg/global/model"
)

const (
	DIRECT_DEVICE   = "direct"   //直连设备
	GATEWAY_DEVICE  = "gateway"  //网关设备
	GATEWAYS_DEVICE = "gatewayS" //网关子设备
	MONITOR_DEVICE  = "monitor"  //监控设备
)

const (
	ATTRIBUTES_TSL = "attributes"
	TELEMETRY_TSL  = "telemetry"
	COMMANDS_TSL   = "commands"
	EVENT_TSL      = "events"
	LOG_TSL        = "logs"
	TAGS_TSL       = "tags"
)

type ProductCategory struct {
	model.BaseAuthModel
	Name        string            `json:"name"  gorm:"type:varchar(128);comment:产品类型名称" validate:"required"`
	Pid         string            `json:"pid" gorm:"type:varchar(64);comment:父产品类型"`
	Path        string            `json:"path" gorm:"type:varchar(255);comment:产品类型路径"`
	Description string            `json:"description"  gorm:"type:varchar(255);comment:产品类型说明"`
	Sort        int64             `json:"sort" gorm:"type:int;comment:排序"`
	Status      string            `gorm:"status;type:varchar(1);comment:状态" json:"status"`
	Children    []ProductCategory `json:"children" gorm:"-"` //子节点
}

type ProductCategoryLabel struct {
	Id       string                 `gorm:"-" json:"id"`
	Name     string                 `gorm:"-" json:"name"`
	Children []ProductCategoryLabel `gorm:"-" json:"children"`
}

type Product struct {
	model.BaseAuthModel
	Name              string `json:"name"  gorm:"type:varchar(128);comment:产品名称" validate:"required"`
	PhotoUrl          string `json:"photoUrl"  gorm:"type:varchar(255);comment:图片地址"`
	Description       string `json:"description"  gorm:"type:varchar(255);comment:产品说明"`
	ProductCategoryId string `json:"productCategoryId" gorm:"type:varchar(64);comment:产品类型Id" validate:"required"`
	ProtocolName      string `json:"protocolName" gorm:"type:varchar(64);comment:协议名称"` //MQTT COAP WebSocket LwM2M
	DeviceType        string `json:"deviceType" gorm:"type:varchar(64);comment:设备类型"`   // 直连设备 网关设备 网关子设备 监控设备
	RuleChainId       string `json:"ruleChainId" gorm:"type:varchar(64);comment:规则链Id"` //可空，如果空就走根规则链
	Status            string `gorm:"type:varchar(1);comment:状态" json:"status"`
}

type ProductRes struct {
	Product
	ProductCategory ProductCategory `json:"productCategory"`
}

type ProductTemplate struct {
	model.BaseModel
	Pid         string `json:"pid" gorm:"type:varchar(64);comment:产品Id" validate:"required"`
	Classify    string `json:"classify" gorm:"type:varchar(64);comment:模型归类" validate:"required"` // 属性 遥测 命令 事件
	Name        string `json:"name"  gorm:"type:varchar(64);comment:名称" validate:"required"`
	Key         string `json:"key"  gorm:"type:varchar(64);comment:标识" validate:"required"`
	Description string `json:"description"  gorm:"type:varchar(255);comment:属性说明"`
	Type        string `json:"type"  gorm:"type:varchar(64);comment:数据类型"` //["int32","float","double","array","bool","enum","date","struct","string"]
	Define      Define `json:"define"  gorm:"type:json;comment:数据约束"`
}

type ProductOta struct {
	model.BaseModel
	Pid         string `json:"pid" gorm:"comment:产品Id" validate:"required"`
	Name        string `json:"name"  gorm:"type:varchar(64);comment:固件名称" validate:"required"`
	Version     string `json:"version" gorm:"type:varchar(64);comment:固件版本" validate:"required"`
	IsLatest    bool   `json:"isLatest" gorm:"comment:是最新固件"`
	Url         string `json:"url" gorm:"type:varchar(128);comment:下载地址" validate:"required"`
	Check       string `json:"check" gorm:"type:varchar(128);comment:md5校验值"`
	Description string `json:"description"  gorm:"type:varchar(255);comment:说明"`
}

type Define map[string]any

func (a Define) Value() (driver.Value, error) {
	return json.Marshal(a)
}
func (a *Define) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
