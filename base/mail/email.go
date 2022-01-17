package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

type Mail struct {
	Host     string `json:"host"`     // 服务器地址
	Port     int    `json:"port"`     // 服务器端口
	From     string `json:"from"`     // 邮箱账号
	Nickname string `json:"nickname"` // 发件人
	Secret   string `json:"secret"`   // 邮箱密码
	IsSSL    bool   `json:"isSsl"`    // 是否开启ssl
}

func (m Mail) Email(to, subject string, body string) error {
	tos := strings.Split(to, ",")
	return m.send(tos, subject, body)
}

//@function: ErrorToEmail
//@description: 给email中间件错误发送邮件到指定邮箱
//@param: subject string, body string
//@return: error

func (m Mail) ErrorToEmail(to, subject string, body string) error {
	tos := strings.Split(to, ",")
	if tos[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		tos = tos[:len(tos)-1]
	}
	return m.send(tos, subject, body)
}

//@function: send
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func (m Mail) send(to []string, subject string, body string) error {

	auth := smtp.PlainAuth("", m.From, m.Secret, m.Host)
	e := email.NewEmail()
	if m.Nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", m.Nickname, m.From)
	} else {
		e.From = m.From
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", m.Host, m.Port)
	if m.IsSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: m.Host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
