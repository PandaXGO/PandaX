package model

import (
	"time"
)

type BaseModel struct {
	Id        string    `json:"id" gorm:"primary_key;"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime" form:"create_time"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime" form:"update_time"`
}

type BaseAuthModel struct {
	Id        string    `json:"id" gorm:"primary_key;"`
	Owner     string    `json:"owner"  gorm:"type:varchar(64);comment:创建者,所有者"`
	OrgId     int64     `json:"orgId"  gorm:"type:int;comment:机构ID"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime" form:"create_time"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime" form:"update_time"`
}
