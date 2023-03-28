package nodes

import (
	"dz-iot-server/rule_engine/message"
	"github.com/sirupsen/logrus"
)

type externalDingNode struct {
	bareNode
	WebHook   string   `json:"webHook" yaml:"webHook"`
	MsgType   string   `json:"msgType" yaml:"msgType"`
	Content   string   `json:"content" yaml:"content"`
	IsAtAll   bool     `json:"isAtAll" yaml:"isAtAll"`
	atMobiles []string `json:"atMobiles" yaml:"atMobiles"`
}

type externalDingNodeFactory struct{}

func (f externalDingNodeFactory) Name() string     { return "ExternalDingNode" }
func (f externalDingNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalDingNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalDingNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalDingNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalDingNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	return nil
}
