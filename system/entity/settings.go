package entity

import (
	"encoding/json"
	"pandax/base/model"
)

type SysSettings struct {
	model.BaseAutoModel
	Key     string          `gorm:"column:key; type:varchar(64)" json:"key" form:"key"`
	Content json.RawMessage `gorm:"column:content; type:json" json:"content" form:"content"` // 配置内容
}
