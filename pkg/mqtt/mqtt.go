package mqtt

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"time"
)

const DefaultDownStreamClientId = `@panda.iothub.internal.clientId`

type IothubMqttClient struct {
	Client mqtt.Client
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	logrus.Infof("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	logrus.Infof("Connect lost: %v", err)
}

func InitMqtt(broker, username, password string) *IothubMqttClient {
	server := fmt.Sprintf("tcp://%s", broker)
	client := GetMqttClinent(server, username, password)
	return &IothubMqttClient{
		Client: client,
	}
}

func GetMqttClinent(server, username, password string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(server)
	time.Now().Unix()
	// 默认下行ID iothub会过滤掉改Id
	opts.SetClientID(DefaultDownStreamClientId)
	if username != "" {
		opts.SetUsername(username)
	}
	if password != "" {
		opts.SetPassword(password)
	}
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

// 订阅
func (imc *IothubMqttClient) Sub(topic string, qos byte, handler mqtt.MessageHandler) {
	if token := imc.Client.Subscribe(topic, qos, handler); token.Wait() && token.Error() != nil {
		logrus.Infof(token.Error().Error())
	}
}

// 取消订阅
func (imc *IothubMqttClient) UnSub(topic string) {
	if token := imc.Client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		logrus.Infof(token.Error().Error())
	}
}

// 发布
func (imc *IothubMqttClient) Pub(topic string, qos byte, payload string) error {
	token := imc.Client.Publish(topic, qos, false, payload)
	if token.WaitTimeout(2*time.Second) == false {
		return errors.New("推送消息超时")
	}
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}
