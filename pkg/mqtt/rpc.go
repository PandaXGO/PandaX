package mqtt

import (
	"encoding/json"
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
	Params string `json:"params"`
}

func (rpc RpcRequest) RequestCmd(rpcPayload RpcPayload) error {
	topic := fmt.Sprintf(RpcReqTopic, rpc.RequestId)
	payload, err := json.Marshal(rpcPayload)
	if err != nil {
		return err
	}
	err = rpc.Client.Pub(topic, 0, string(payload))
	if err != nil {
		return err
	}
	if rpc.Mode == "single" {
		return nil
	}
	respTopic := fmt.Sprintf(RpcRespTopic, rpc.RequestId)
	rpc.Client.Sub(respTopic, 0, mqMessagePubHandler)
	if rpc.Timeout == 0 {
		rpc.Timeout = 30
	}
	go func() {
		select {
		case <-time.After(time.Duration(rpc.Timeout) * time.Second):
			rpc.Client.UnSub(respTopic)
		}
	}()
	return nil
}

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
var mqMessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//log.Println(fmt.Sprintf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic()))
}

// RespondTpc 处理设备端请求服务端方法
func (rpc RpcRequest) RespondTpc(reqPayload RpcPayload) error {
	topic := fmt.Sprintf(RpcRespTopic, rpc.RequestId)
	if reqPayload.Params == "getCurrentTime" {
		unix := time.Now().Unix()
		msg := fmt.Sprintf("%d", unix)
		return rpc.Client.Pub(topic, 0, msg)
	}

	return nil
}

func (rpc *RpcRequest) GetRequestId() {
	rand.Seed(time.Now().UnixNano())
	// 生成随机整数
	rpc.RequestId = rand.Intn(10000) + 1 // 生成0到99之间的随机整数
}
