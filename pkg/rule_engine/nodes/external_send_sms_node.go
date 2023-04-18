package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type externalSendSmsNode struct {
	bareNode
	SecretId      string                 `json:"secretId" yaml:"secretId"`
	SecretKey     string                 `json:"secretKey" yaml:"secretKey"`
	SdkAppId      string                 `json:"sdkAppId" yaml:"sdkAppId"`           //应用Id（腾讯） 或 签名名称（阿里）
	PhoneNumber   string                 `json:"phoneNumber" yaml:"phoneNumber"`     //发送到手机号
	TemplateId    string                 `json:"templateId" yaml:"templateId"`       //短信模板Id
	TemplateParam map[string]interface{} `json:"templateParam" yaml:"templateParam"` //模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空*/
}

type externalSendSmsNodeFactory struct{}

func (f externalSendSmsNodeFactory) Name() string     { return "SendSmsNode" }
func (f externalSendSmsNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalSendSmsNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalSendSmsNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalSendSmsNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalSendSmsNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	//failureLabelNode := n.GetLinkedNode("Failure")

	return successLabelNode.Handle(msg)
}
