package nodes

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"pandax/iothub/client/mqttclient"
	"pandax/iothub/client/tcpclient"
	"pandax/pkg/global"
	"pandax/pkg/global_model"
	"pandax/pkg/rule_engine/message"
)

type rpcRequestToDeviceNode struct {
	bareNode
	Timeout int `json:"timeout"`
}

type rpcRequestToDeviceNodeFactory struct{}

func (f rpcRequestToDeviceNodeFactory) Name() string     { return "RpcRequestToDeviceNode" }
func (f rpcRequestToDeviceNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f rpcRequestToDeviceNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f rpcRequestToDeviceNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &rpcRequestToDeviceNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *rpcRequestToDeviceNode) Handle(msg *message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.MsgType)
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")
	if msg.Msg.GetValue("method") == nil || msg.Msg.GetValue("params") == nil {
		return errors.New("指令下发格式错误")
	}
	var datas = global_model.RpcPayload{
		Method: msg.Msg.GetValue("method").(string),
		Params: msg.Msg.GetValue("params"),
	}
	payload, _ := json.Marshal(datas)
	mode := mqttclient.SingleMode
	if n.Timeout > 0 {
		mode = mqttclient.DoubleMode
	}
	// 判断设备协议，根据不通协议，发送不通内容
	deviceProtocol := global.MQTTProtocol
	if msg.Metadata.GetValue("deviceProtocol") != nil && msg.Metadata.GetValue("deviceProtocol").(string) != "" {
		deviceProtocol = msg.Metadata.GetValue("deviceProtocol").(string)
	}
	var err error
	deviceId := msg.Metadata.GetValue("deviceId").(string)
	if deviceProtocol == global.MQTTProtocol {
		var rpc = &mqttclient.RpcRequest{Mode: mode, Timeout: n.Timeout}
		rpc.GetRequestId()
		err = rpc.RequestCmd(deviceId, string(payload))
	}
	if deviceProtocol == global.TCPProtocol {
		err = tcpclient.Send(deviceId, string(payload))
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
