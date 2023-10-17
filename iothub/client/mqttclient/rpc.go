package mqttclient

import (
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

const (
	SingleMode = "single"
	DoubleMode = "double"
)

type RpcRequest struct {
	Client    *IothubMqttClient
	RequestId int
	Mode      string //单向、双向 单项只发送不等待响应  双向需要等到响应
	Timeout   int    // 设置双向时，等待的超时时间
}

// RequestCmd 下发指令
func (rpc RpcRequest) RequestCmd(rpcPayload string) (respPayload string, err error) {
	topic := fmt.Sprintf(RpcReqTopic, rpc.RequestId)
	err = rpc.Client.Pub(topic, 0, rpcPayload)
	if err != nil {
		return "", err
	}
	if rpc.Mode == "" || rpc.Mode == SingleMode {
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
		rpc.Timeout = 20
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

func (rpc RpcRequest) Pub(reqPayload string) error {
	topic := fmt.Sprintf(RpcRespTopic, rpc.RequestId)
	return rpc.Client.Pub(topic, 0, reqPayload)
}

func (rpc *RpcRequest) GetRequestId() {
	rand.Seed(time.Now().UnixNano())
	// 生成随机整数
	rpc.RequestId = rand.Intn(10000) + 1 // 生成0到99之间的随机整数
}
