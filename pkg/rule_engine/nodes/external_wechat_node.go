package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type externalWechatNode struct {
	bareNode
	WebHook   string   `json:"webHook" yaml:"webHook"`
	MsgType   string   `json:"msgType" yaml:"msgType"`
	Content   string   `json:"content" yaml:"content"`
	IsAtAll   bool     `json:"isAtAll" yaml:"isAtAll"`
	atMobiles []string `json:"atMobiles" yaml:"atMobiles"`
}

type externalWechatNodeFactory struct{}

func (f externalWechatNodeFactory) Name() string     { return "ExternalWechatNode" }
func (f externalWechatNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalWechatNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalWechatNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalWechatNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalWechatNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	return nil
}
