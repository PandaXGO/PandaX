package entity

import "github.com/XM-GO/PandaKit/model"

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/13 14:47
 **/
type ResEmail struct {
	MailId   int64  `json:"mailId" gorm:"primaryKey;AUTO_INCREMENT;comment:主键编码"`
	Category string `json:"category" grom:"type:varchar(1);comment:分类qq或163"` // 0 163邮箱，1 qq邮箱 2 企业邮箱
	Host     string `json:"host" grom:"type:varchar(64);comment:服务器地址"`       // 服务器地址
	Port     int    `json:"port" grom:"type:int;comment:服务器端口"`               // 服务器端口
	From     string `json:"from" grom:"type:varchar(64);comment:邮箱账号"`        // 邮箱账号
	Nickname string `json:"nickname" grom:"type:varchar(64);comment:发件人"`     // 发件人
	Secret   string `json:"secret" grom:"type:varchar(64);comment:邮箱密码"`      // 邮箱密码
	IsSSL    bool   `json:"isSsl" grom:"comment:是否开启ssl"`                     // 是否开启ssl
	Status   string `json:"status" grom:"type:varchar(1);comment:状态"`
	model.BaseModel
}
