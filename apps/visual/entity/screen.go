package entity

import "github.com/XM-GO/PandaKit/model"

type VisualScreenGroup struct {
	Id       int64               `gorm:"id;primary_key;type:varchar(64);comment:Id" json:"id"`
	Name     string              `gorm:"name;type:varchar(64);comment:分组名称" json:"name"`
	Pid      int64               `gorm:"pid;type:varchar(64);comment:父Id" json:"pid"`
	Path     string              `gorm:"path;type:varchar(64);comment:路径" json:"path"`
	Sort     int64               `gorm:"sort;type:int;comment:排序" json:"sort"`
	Status   string              `gorm:"status;type:varchar(1);comment:状态" json:"status"`
	Children []VisualScreenGroup `json:"children" gorm:"-"`
}

type ScreenGroupLabel struct {
	Id       int64              `gorm:"-" json:"id"`
	Name     string             `gorm:"-" json:"name"`
	Children []ScreenGroupLabel `gorm:"-" json:"children"`
}

func (VisualScreenGroup) TableName() string {
	return "visual_screen_group"
}

type VisualScreen struct {
	UserId         int64  `gorm:"userId;type:int;comment:用户Id" json:"userId"`
	ScreenId       string `gorm:"primary_key;" json:"screenId"`
	GroupId        int64  `gorm:"screenGroup;type:int;comment:分组Id" json:"groupId"`
	ScreenName     string `gorm:"screenName;type:varchar(50);comment:名称" json:"screenName"`
	ScreenDataJson string `gorm:"screenDataJson;type:text;comment:Json数据" json:"screenDataJson"`
	ScreenBase64   string `gorm:"screenBase64;type:text;comment:Base64缩略图" json:"screenBase64"` //缩略图 base64
	ScreenRemark   string `gorm:"screenRemark;type:varchar(255);comment:说明" json:"screenRemark"`
	Status         string `gorm:"status;type:varchar(1);comment:状态" json:"status"`
	Creator        string `json:"creator"` //创建者
	model.BaseModel
}

func (VisualScreen) TableName() string {
	return "visual_screen"
}
