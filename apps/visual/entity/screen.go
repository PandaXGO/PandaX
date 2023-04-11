package entity

import "github.com/XM-GO/PandaKit/model"

type VisualScreenGroup struct {
	Id    string `gorm:"id;primary_key;type:varchar(64);comment:Id" json:"id"`
	Name  string `gorm:"name;type:varchar(64);comment:分组名称" json:"name"`
	Pid   string `gorm:"pid;type:varchar(64);comment:父Id" json:"pid"`
	Level int64  `gorm:"level;type:int;comment:等级" json:"level"`
}

func (VisualScreenGroup) TableName() string {
	return "visual_screen_group"
}

type VisualScreen struct {
	UserId         string `gorm:"userId;type:varchar(64);comment:用户Id" json:"userId"`
	ScreenId       string `gorm:"primary_key;" json:"screenId"`
	ScreenName     string `gorm:"screenName;type:varchar(50);comment:名称" json:"screenName"`
	ScreenDataJson string `gorm:"screenDataJson;type:varchar(50);comment:Json数据" json:"screenDataJson"`
	ScreenBase64   string `gorm:"screenBase64;type:varchar(50);comment:Base64缩略图" json:"screenBase64"` //缩略图 base64
	ScreenRemark   string `gorm:"screenRemark;type:varchar(50);comment:说明" json:"screenRemark"`
	Status         string `gorm:"status;type:varchar(50);comment:状态" json:"status"`
	Creator        string `json:"creator"` //创建者
	model.BaseModel
}

func (VisualScreen) TableName() string {
	return "visual_screen"
}
