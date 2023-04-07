package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type externalKafkaNode struct {
	bareNode
	Server   string `json:"server" yaml:"server"`
	Topic    string `json:"topic" yaml:"topic"`
	KafkaCli string
}

type externalKafkaNodeFactory struct{}

func (f externalKafkaNodeFactory) Name() string     { return "KafkaNode" }
func (f externalKafkaNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalKafkaNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalKafkaNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalKafkaNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalKafkaNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	return nil
}
