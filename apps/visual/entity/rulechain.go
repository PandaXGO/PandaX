package entity

import (
	"github.com/XM-GO/PandaKit/model"
	"time"
)

type VisualRuleChainBase struct {
	UserId     string `gorm:"userId;type:varchar(64);comment:用户Id" json:"userId"`
	RuleId     string `gorm:"primary_key;" json:"ruleId"`
	Root       bool   `gorm:"root;comment:根链" json:"root"`
	RuleName   string `gorm:"ruleName;type:varchar(50);comment:名称" json:"ruleName"`
	RuleBase64 string `gorm:"ruleBase64;type:longtext;comment:Base64缩略图" json:"ruleBase64"` //缩略图 base64
	RuleRemark string `gorm:"ruleRemark;type:varchar(256);comment:说明" json:"ruleRemark"`
	Status     string `gorm:"status;type:varchar(1);comment:状态" json:"status"`
	Creator    string `json:"creator"` //创建者
	model.BaseModel
}

type VisualRuleChain struct {
	VisualRuleChainBase
	RuleDataJson string `gorm:"ruleDataJson;type:longtext;comment:Json数据" json:"ruleDataJson"`
}

func (VisualRuleChain) TableName() string {
	return "visual_rule_chain"
}

type VisualRuleChainMsgLog struct {
	MessageId  string    `json:"message_id"`
	MsgType    string    `json:"msg_type"`
	DeviceName string    `json:"device_name"`
	Ts         time.Time `json:"ts"`
	Content    string    `json:"content"`
	CreatedAt  time.Time // 创建时间
}

func (VisualRuleChainMsgLog) TableName() string {
	return "visual_rule_chain_msg_log"
}
