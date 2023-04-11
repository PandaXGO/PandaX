package entity

import (
	"github.com/XM-GO/PandaKit/model"
)

type VisualRuleChain struct {
	UserId       string `gorm:"userId;type:varchar(64);comment:用户Id" json:"userId"`
	RuleId       string `gorm:"primary_key;" json:"ruleId"`
	RuleName     string `gorm:"ruleName;type:varchar(50);comment:名称" json:"ruleName"`
	RuleDataJson string `gorm:"ruleDataJson;type:text;comment:Json数据" json:"ruleDataJson"`
	RuleBase64   string `gorm:"ruleBase64;type:text;comment:Base64缩略图" json:"ruleBase64"` //缩略图 base64
	RuleRemark   string `gorm:"ruleRemark;type:varchar(256);comment:说明" json:"ruleRemark"`
	Status       string `gorm:"status;type:varchar(1);comment:状态" json:"status"`
	Creator      string `json:"creator"` //创建者
	model.BaseModel
}

func (VisualRuleChain) TableName() string {
	return "visual_rule_chain"
}
