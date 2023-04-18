package nodes

import (
	"encoding/json"
	"fmt"
	"pandax/pkg/rule_engine/message"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

type externalMqttNode struct {
	bareNode
	TopicPattern      string `json:"topicPattern"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	Host              string `json:"host"`
	Port              string `json:"port"`
	ConnectTimeoutSec int    `json:"connectTimeoutSec"`
	ClientId          string `json:"clientId"`
	CleanSession      bool   `json:"cleanSession"`
	Ssl               bool   `json:"ssl"`
	MqttCli           mqtt.Client
}

type externalMqttNodeFactory struct{}

func (f externalMqttNodeFactory) Name() string     { return "MqttNode" }
func (f externalMqttNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalMqttNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalMqttNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalMqttNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	_, err := decodePath(meta, node)
	if err != nil {
		return node, err
	}
	broker := fmt.Sprintf("tcp://%s:%s", node.Host, node.Port)
	opts := mqtt.NewClientOptions().AddBroker(broker)
	if node.ClientId != "" {
		opts.SetClientID(node.ClientId)
	}
	opts.SetCleanSession(node.CleanSession)
	opts.SetConnectTimeout(time.Duration(node.ConnectTimeoutSec) * time.Second)
	if node.Username != "" {
		opts.SetUsername(node.Username)
	}
	if node.Password != "" {
		opts.SetPassword(node.Password)
	}
	node.MqttCli = mqtt.NewClient(opts)

	if token := node.MqttCli.Connect(); token.Wait() && token.Error() != nil {
		logrus.WithError(token.Error())
		return nil, token.Error()
	}
	return node, nil
}

func (n *externalMqttNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	defer n.MqttCli.Disconnect(1000)
	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	topic := n.TopicPattern //need fix add msg.metadata in it
	sendmqttmsg, err := json.Marshal(msg.GetMsg())
	if err != nil {
		return err
	}
	token := n.MqttCli.Publish(topic, 1, false, sendmqttmsg)
	if token.Wait() && token.Error() != nil {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return token.Error()
		}
	}
	if successLabelNode != nil {
		return successLabelNode.Handle(msg)
	}
	return nil
}
