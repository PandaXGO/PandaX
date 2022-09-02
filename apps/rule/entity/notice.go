package entity

import (
	"encoding/json"
	"github.com/XM-GO/PandaKit/model"
)

type RuleNotice struct {
	model.BaseAutoModel
	UserId      string          `json:"user_id"`
	Name        string          `json:"name"`
	Category    string          `json:"category"` // http，mail
	Description string          `json:"description"`
	Model       string          `json:"model"` // 配置或模板   setting or template
	ExParam     json.RawMessage `json:"ex_param" gorm:"type:jsonb;comment: 拓展参数"`
}

func (RuleNotice) TableName() string {
	return "rule_notice"
}

type RestSetting struct {
	Method             string            `json:"method"`
	Url                string            `json:"url"`
	Headers            map[string]string `json:"headers"`
	BodyType           string            `json:"bodyType"`
	Timeout            int64             `json:"timeout"`
	CertificationPath  string            `json:"certificationPath"`
	PrivateKeyPath     string            `json:"privateKeyPath"`
	RootCaPath         string            `json:"rootCaPath"`
	InsecureSkipVerify bool              `json:"insecureSkipVerify"`
}

type MqttSetting struct {
	Server             string `json:"server"`
	Topic              string `json:"topic"`
	Qos                byte   `json:"qos"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Retained           bool   `json:"retained"`
	CertificationPath  string `json:"certificationPath"`
	PrivateKeyPath     string `json:"privateKeyPath"`
	RootCaPath         string `json:"rootCaPath"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify"`
}

type MailSetting struct {
	Host     string `json:"host"`     // 服务器地址
	Port     int    `json:"port"`     // 服务器端口
	From     string `json:"from"`     // 邮箱账号
	Nickname string `json:"nickname"` // 发件人
	Secret   string `json:"secret"`   // 邮箱密码
	IsSSL    bool   `json:"isSsl"`    // 是否开启ssl
}

type ScriptSetting struct {
	Category string `json:"category"` // 0 python，1 javascript 2 shell
}

type DataTemplate struct {
	DataTemplate string `json:"dataTemplate"`
}
