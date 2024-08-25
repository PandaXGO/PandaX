package entity

import (
	"github.com/PandaXGO/PandaKit/model"
	"time"
)

type LogLogin struct {
	InfoId        int64     `json:"infoId" gorm:"primary_key;AUTO_INCREMENT"` //主键
	Username      string    `json:"username" gorm:"type:varchar(128);comment:用户名"`
	Status        string    `json:"status" gorm:"type:varchar(1);comment:状态"`
	Ipaddr        string    `json:"ipaddr" gorm:"type:varchar(255);comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"type:varchar(255);comment:归属地"`
	Browser       string    `json:"browser" gorm:"type:varchar(255);comment:浏览器"`
	Os            string    `json:"os" gorm:"type:varchar(255);comment:系统"`
	Platform      string    `json:"platform" gorm:"type:varchar(255);comment:固件"`
	LoginTime     time.Time `json:"loginTime" gorm:"type:timestamp;comment:登录时间"`
	CreateBy      string    `json:"createBy" gorm:"type:varchar(128);comment:创建人"`
	UpdateBy      string    `json:"updateBy" gorm:"type:varchar(128);comment:更新者"`
	Params        string    `json:"params" gorm:"-"`
	Remark        string    `json:"remark" gorm:"type:varchar(255);"` //备注
	Msg           string    `json:"msg" gorm:"type:varchar(255);"`
	model.BaseModel
}
