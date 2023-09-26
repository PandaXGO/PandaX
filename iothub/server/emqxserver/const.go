package emqxserver

import (
	"pandax/pkg/rule_engine/message"
	"strings"
)

const (
	// Topic 上行
	RowTopic        = `v1/devices/me/row`
	TelemetryTopic  = `v1/devices/me/telemetry`
	AttributesTopic = `v1/devices/me/attributes`

	//网关子设备
	AttributesGatewayTopic = "v1/gateway/attributes"
	TelemetryGatewayTopic  = "v1/gateway/telemetry"
	ConnectGatewayTopic    = "v1/gateway/connect"
	DisconnectGatewayTopic = "v1/gateway/disconnect"

	RpcReqReg = `v1/devices/me/rpc/request/(.*?)$`
)

var IotHubTopic = NewIotHubTopic()

type TopicMeg map[string]string

// 消息的来源类型
func NewIotHubTopic() TopicMeg {
	return map[string]string{
		AttributesTopic:        message.AttributesMes,
		RowTopic:               message.RowMes,
		TelemetryTopic:         message.TelemetryMes,
		AttributesGatewayTopic: message.GATEWAY,
		TelemetryGatewayTopic:  message.GATEWAY,
		ConnectGatewayTopic:    message.GATEWAY,
		DisconnectGatewayTopic: message.GATEWAY,
	}
}

func (iht TopicMeg) GetMessageType(topic string) string {
	if meg, ok := iht[topic]; ok {
		return meg
	}
	if strings.Contains(topic, "v1/devices/me/rpc/request") {
		return message.RpcRequestMes
	}
	return ""
}
