package nodes

import (
	"encoding/json"
	"errors"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/iothub/client/mqttclient"
	"pandax/iothub/client/tcpclient"
	"pandax/iothub/client/udpclient"
	"pandax/kit/utils"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"pandax/pkg/rule_engine/message"
	"time"

	"github.com/kakuilan/kgo"
)

type rpcRequestToDeviceNode struct {
	bareNode
	Timeout int `json:"timeout"`
}

type rpcRequestToDeviceNodeFactory struct{}

func (f rpcRequestToDeviceNodeFactory) Name() string     { return "RpcRequestToDeviceNode" }
func (f rpcRequestToDeviceNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f rpcRequestToDeviceNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f rpcRequestToDeviceNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &rpcRequestToDeviceNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *rpcRequestToDeviceNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")
	if msg.Msg.GetValue("method") == nil || msg.Msg.GetValue("params") == nil {
		return errors.New("指令下发格式错误")
	}
	deviceId := msg.Metadata.GetValue("deviceId").(string)
	// 创建请求格式
	var datas = model.RpcPayload{
		Method: msg.Msg.GetValue("method").(string),
		Params: msg.Msg.GetValue("params"),
	}
	payload, _ := json.Marshal(datas)

	// 构建指令记录
	var data entity.DeviceCmdLog
	data.Id = utils.GenerateID()
	data.DeviceId = deviceId
	data.CmdName = datas.Method
	data.CmdContent = kgo.KConv.ToStr(datas.Params)
	data.Mode = msg.Metadata.GetValue("mode").(string)
	data.State = "2"
	data.RequestTime = time.Now().Format("2006-01-02 15:04:05")

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
	if deviceProtocol == global.MQTTProtocol || deviceProtocol == global.CoAPProtocol || deviceProtocol == global.LwM2MProtocol {
		var rpc = &mqttclient.RpcRequest{Mode: mode, Timeout: n.Timeout, RequestId: data.Id}
		err = rpc.RequestCmd(deviceId, string(payload))
	}
	if deviceProtocol == global.TCPProtocol {
		err = tcpclient.Send(deviceId, string(payload))
	}
	if deviceProtocol == global.UDPProtocol {
		err = udpclient.Send(deviceId, string(payload))
	}

	if err != nil {
		data.State = "1"
		services.DeviceCmdLogModelDao.Insert(data)
		n.Debug(msg, message.DEBUGOUT, err.Error())
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}

	if successLableNode != nil {
		data.State = "0"
		services.DeviceCmdLogModelDao.Insert(data)
		n.Debug(msg, message.DEBUGOUT, "")
		return successLableNode.Handle(msg)
	}
	return nil
}
