package mqttclient

import (
	"encoding/json"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"time"
)

const (
	RpcRespTopic = `v1/devices/me/rpc/response/%d`
	RpcReqTopic  = `v1/devices/me/rpc/request/%d`
)

type RpcRequest struct {
	Client    *IothubMqttClient
	RequestId int
	Mode      string //单向、双向 单项只发送不等待响应  双向需要等到响应
	Timeout   int    // 设置双向时，等待的超时时间
}

type RpcPayload struct {
	Method string `json:"method"`
	Params any    `json:"params"`
}

// RequestCmd 下发指令
func (rpc RpcRequest) RequestCmd(rpcPayload RpcPayload) (respPayload string, err error) {
	topic := fmt.Sprintf(RpcReqTopic, rpc.RequestId)
	payload, err := json.Marshal(rpcPayload)
	if err != nil {
		return "", err
	}
	err = rpc.Client.Pub(topic, 0, string(payload))
	if err != nil {
		return "", err
	}
	if rpc.Mode == "single" {
		return "", nil
	}
	// 双向才会执行
	repsChan := make(chan string)
	respTopic := fmt.Sprintf(RpcRespTopic, rpc.RequestId)
	// 订阅回调
	mqMessagePubHandler := func(client mqtt.Client, msg mqtt.Message) {
		if repsChan != nil {
			repsChan <- string(msg.Payload())
		}
	}
	rpc.Client.Sub(respTopic, 0, mqMessagePubHandler)
	if rpc.Timeout == 0 {
		rpc.Timeout = 30
	}
	defer func() {
		close(repsChan)
		rpc.Client.UnSub(respTopic)
	}()
	for {
		select {
		case <-time.After(time.Duration(rpc.Timeout) * time.Second):
			return "", errors.New("设备指令响应超时")
		case resp := <-repsChan:
			return resp, nil
		}
	}
}

// RequestAttributes rpc 下发属性
func (rpc RpcRequest) RequestAttributes(rpcPayload RpcPayload) error {
	topic := fmt.Sprintf(RpcReqTopic, rpc.RequestId)
	if rpcPayload.Method == "" {
		rpcPayload.Method = "setAttributes"
	}
	payload, err := json.Marshal(rpcPayload)
	if err != nil {
		return err
	}
	return rpc.Client.Pub(topic, 0, string(payload))
}

// 响应数据处理
/*var mqMessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//log.Println(fmt.Sprintf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic()))
}*/

// RespondTpc 处理设备端请求服务端方法
func (rpc RpcRequest) RespondTpc(reqPayload RpcPayload) error {
	topic := fmt.Sprintf(RpcRespTopic, rpc.RequestId)
	//TODO 此处处理设备的请求参数逻辑
	//自己定义请求逻辑
	if reqPayload.Params == "getCurrentTime" {
		unix := time.Now().Unix()
		msg := fmt.Sprintf("%d", unix)
		return rpc.Client.Pub(topic, 0, msg)
	}
	// 获取属性 ...
	return nil
}

func (rpc *RpcRequest) GetRequestId() {
	rand.Seed(time.Now().UnixNano())
	// 生成随机整数
	rpc.RequestId = rand.Intn(10000) + 1 // 生成0到99之间的随机整数
}
