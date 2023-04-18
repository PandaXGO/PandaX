package nodes

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
	"strings"
	"time"
)

type externalKafkaNode struct {
	bareNode
	Server            string                 `json:"server" yaml:"server"`         //kafka集群 "10.130.138.164:9092,10.130.138.165:9093"
	Topic             string                 `json:"topic" yaml:"topic"`           //topic
	KeyPattern        string                 `json:"keyPattern" yaml:"keyPattern"` //metadataKey or messageKey
	ProducesBatchSize int64                  `json:"producesBatchSize" yaml:"producesBatchSize"`
	OtherProperties   map[string]interface{} `json:"otherProperties"` //发送的其他参数
	KafkaCli          sarama.SyncProducer
}

type externalKafkaNodeFactory struct{}

func (f externalKafkaNodeFactory) Name() string     { return "KafkaNode" }
func (f externalKafkaNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalKafkaNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalKafkaNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalKafkaNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	_, err := decodePath(meta, node)
	if err != nil {
		return node, err
	}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Producer.MaxMessageBytes = int(node.ProducesBatchSize)
	p, err := sarama.NewSyncProducer(strings.Split(node.Server, ","), config)
	if err != nil {
		logrus.Errorf("sarama.NewSyncProducer err, message=%s \n", err)
		return node, err
	}
	node.KafkaCli = p
	return node, nil
}

func (n *externalKafkaNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	defer n.KafkaCli.Close()
	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	value := sarama.ByteEncoder("")
	if n.KeyPattern == "metadataKey" {
		marshal, err := json.Marshal(msg.GetMetadata().GetValues())
		if err != nil {
			return err
		}
		value = marshal
	} else {
		marshal, err := json.Marshal(msg.GetMsg())
		if err != nil {
			return err
		}
		value = marshal
	}

	kafkaM := &sarama.ProducerMessage{
		Topic: n.Topic,
		Value: value,
	}
	_, _, err := n.KafkaCli.SendMessage(kafkaM)
	if err != nil {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLabelNode != nil {
		return successLabelNode.Handle(msg)
	}

	return nil
}
