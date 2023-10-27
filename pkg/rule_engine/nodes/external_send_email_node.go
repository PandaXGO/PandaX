package nodes

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"pandax/pkg/rule_engine/message"
	"strings"

	"github.com/jordan-wright/email"
)

type externalSendEmailNode struct {
	bareNode
	Host     string `json:"host"`     // 服务器地址
	Port     int    `json:"port"`     // 服务器端口
	From     string `json:"from"`     // 邮箱账号
	Nickname string `json:"nickname"` // 发件人
	Secret   string `json:"secret"`   // 邮箱密码
	IsSSL    bool   `json:"isSsl"`    // 是否开启ssl

	To      string `json:"to"`      //收件人
	Subject string `json:"subject"` //主题
	Body    string `json:"body"`    //内容
}

type externalSendEmailNodeFactory struct{}

func (f externalSendEmailNodeFactory) Name() string     { return "SendEmailNode" }
func (f externalSendEmailNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalSendEmailNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalSendEmailNodeFactory) Create(id string, meta Properties) (Node, error) {

	node := &externalSendEmailNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalSendEmailNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")

	tos := strings.Split(n.To, ",")
	if tos[len(tos)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		tos = tos[:len(tos)-1]
	}
	err := n.send(tos, *msg)
	if err != nil {
		n.Debug(msg, message.DEBUGOUT, err.Error())
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLabelNode != nil {
		n.Debug(msg, message.DEBUGOUT, "")
		return successLabelNode.Handle(msg)
	}
	return nil
}

func (m *externalSendEmailNode) send(to []string, msg message.Message) error {

	auth := smtp.PlainAuth("", m.From, m.Secret, m.Host)
	e := email.NewEmail()
	if m.Nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", m.Nickname, m.From)
	} else {
		e.From = m.From
	}
	e.To = to
	e.Subject = m.Subject
	template, err := ParseTemplate(m.Body, msg.Metadata)
	if err != nil {
		return err
	}
	e.HTML = []byte(template)
	hostAddr := fmt.Sprintf("%s:%d", m.Host, m.Port)
	if m.IsSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: m.Host})
	} else {
		err = e.Send(hostAddr, auth)
	}

	return err
}
