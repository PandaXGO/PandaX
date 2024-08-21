package mqttclient

import (
	"errors"
	"fmt"
)

const (
	RpcRespTopic = `v1/devices/me/rpc/response/%s`
	RpcReqTopic  = `v1/devices/me/rpc/request/%s`
)

const (
	SingleMode = "single"
	DoubleMode = "double"
)

type RpcRequest struct {
	RequestId string
	Mode      string //单向、双向 单项只发送不等待响应  双向需要等到响应
	Timeout   int    // 设置双向时，等待的超时时间
}

// RequestCmd 下发指令
func (rpc RpcRequest) RequestCmd(deviceId, rpcPayload string) error {
	topic := fmt.Sprintf(RpcReqTopic, rpc.RequestId)
	value, ok := Session.Load(deviceId)
	if !ok {
		return errors.New("未获取到设备的MQTT连接")
	}
	return Publish(topic, value.(string), rpcPayload)
}

func (rpc RpcRequest) Pub(deviceId, reqPayload string) error {
	topic := fmt.Sprintf(RpcRespTopic, rpc.RequestId)
	value, ok := Session.Load(deviceId)
	if !ok {
		return errors.New("未获取到设备的MQTT连接")
	}
	return Publish(topic, value.(string), reqPayload)
}
