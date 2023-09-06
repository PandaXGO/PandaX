package entity

import (
	"pandax/pkg/global"
	"time"
)

type RuleChainBaseLabel struct {
	Id       string `json:"id"`
	Root     string `json:"root"`
	RuleName string `json:"ruleName"`
}

type RuleChainBase struct {
	global.BaseAuthModel
	Root       string `json:"root" gorm:"comment:是否根节点,1 根链 0 普通链"`
	RuleName   string `gorm:"ruleName;type:varchar(50);comment:名称" json:"ruleName"`
	RuleBase64 string `gorm:"ruleBase64;type:longtext;comment:Base64缩略图" json:"ruleBase64"` //缩略图 base64
	RuleRemark string `gorm:"ruleRemark;type:varchar(256);comment:说明" json:"ruleRemark"`
}

type RuleChain struct {
	RuleChainBase
	RuleDataJson string `gorm:"ruleDataJson;type:longtext;comment:Json数据" json:"ruleDataJson"`
}

func (RuleChain) TableName() string {
	return "rule_chain"
}

type RuleChainMsgLog struct {
	MessageId  string    `gorm:"message_id;type:varchar(64);comment:消息Id" json:"messageId"`
	MsgType    string    `gorm:"msg_type;type:varchar(64);comment:消息类型" json:"msgType"`
	DeviceId   string    `gorm:"device_id;type:varchar(64);comment:设备ID" json:"deviceId"`
	DeviceName string    `gorm:"device_name;type:varchar(255);comment:设备名称" json:"deviceName"`
	Ts         time.Time `gorm:"ts;type:varchar(64);comment:时间" json:"ts"`
	Content    string    `gorm:"content;type:varchar(1024);comment:内容" json:"content"`
	CreatedAt  time.Time `gorm:"column:create_at" json:"create_at"` // 创建时间
}

func (RuleChainMsgLog) TableName() string {
	return "rule_chain_msg_log"
}
