package nodes

import (
	"errors"
	"github.com/sirupsen/logrus"
	"pandax/iothub/client/mqttclient"
	"pandax/iothub/client/tcpclient"
	"pandax/pkg/global"
	"pandax/pkg/global_model"
	"pandax/pkg/rule_engine/message"
)

type rpcRequestFromDeviceNode struct {
	bareNode
	RequestId int `json:"requestId"`
}

type rpcRequestFromDeviceFactory struct{}

func (f rpcRequestFromDeviceFactory) Name() string     { return "RpcRequestFromDeviceNode" }
func (f rpcRequestFromDeviceFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f rpcRequestFromDeviceFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f rpcRequestFromDeviceFactory) Create(id string, meta Metadata) (Node, error) {
	node := &rpcRequestFromDeviceNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *rpcRequestFromDeviceNode) Handle(msg *message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.MsgType)
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")

	if msg.Msg.GetValue("method") == nil || msg.Msg.GetValue("params") == nil {
		return errors.New("指令请求格式错误")
	}
	var rpcp = global_model.RpcPayload{
		Method: msg.Msg.GetValue("method").(string),
		Params: msg.Msg.GetValue("params"),
	}
	result, err := rpcp.GetRequestResult()
	if err != nil {
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}
	// 判断设备协议，根据不通协议，发送不通内容
	deviceProtocol := global.MQTTProtocol
	if msg.Metadata.GetValue("deviceProtocol") != nil && msg.Metadata.GetValue("deviceProtocol").(string) != "" {
		deviceProtocol = msg.Metadata.GetValue("deviceProtocol").(string)
	}
	deviceId := msg.Metadata.GetValue("deviceId").(string)
	if deviceProtocol == global.MQTTProtocol {
		rpc := &mqttclient.RpcRequest{}
		RequestId := n.RequestId
		if RequestId == 0 {
			if msg.Metadata.GetValue("requestId") == nil {
				rpc.GetRequestId()
			} else {
				RequestId = int(msg.Metadata.GetValue("requestId").(float64))
			}
		} else {
			rpc.RequestId = RequestId
		}
		err = rpc.Pub(deviceId, result)
	}
	if deviceProtocol == global.TCPProtocol {
		err = tcpclient.Send(deviceId, result)
	}
	if err != nil {
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLableNode != nil {
		return successLableNode.Handle(msg)
	}
	return nil
}
