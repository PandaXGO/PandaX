package entity

import (
	"github.com/XM-GO/PandaKit/model"
)

type VisualRuleChain struct {
	UserId       string `json:"userId"`
	RuleId       string `json:"ruleId"`
	RuleName     string `json:"ruleName"`
	RuleDataJson string `json:"ruleDataJson"`
	RuleBase64   string `json:"ruleBase64"` //缩略图 base64
	RuleRemark   string `json:"ruleRemark"`
	Status       string `json:"status"`
	DeviceId     string `json:"deviceId"`
	Creator      string `json:"creator"` //创建者
	model.BaseModel
}
