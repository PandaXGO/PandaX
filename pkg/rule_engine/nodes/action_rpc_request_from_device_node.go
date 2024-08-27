package nodes

import (
	"errors"
	"pandax/apps/device/services"
	"pandax/iothub/client/mqttclient"
	"pandax/iothub/client/tcpclient"
	"pandax/iothub/client/udpclient"
	"pandax/kit/utils"
	devicerpc "pandax/pkg/device_rpc"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
	"time"

	"github.com/kakuilan/kgo"
)

type rpcRequestFromDeviceNode struct {
	bareNode
	RequestId string `json:"requestId"`
}

type rpcRequestFromDeviceFactory struct{}

func (f rpcRequestFromDeviceFactory) Name() string     { return "RpcRequestFromDeviceNode" }
func (f rpcRequestFromDeviceFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f rpcRequestFromDeviceFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f rpcRequestFromDeviceFactory) Create(id string, meta Properties) (Node, error) {
	node := &rpcRequestFromDeviceNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *rpcRequestFromDeviceNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")

	if msg.Msg.GetValue("method") == nil || msg.Msg.GetValue("params") == nil {
		return errors.New("指令请求格式错误")
	}

	var deviceId string
	if did, ok := msg.Metadata.GetValue("deviceId").(string); ok {
		deviceId = did
	} else {
		return errors.New("元数据中为获取到设备ID")
	}
	var rpcp = devicerpc.RpcPayload{
		Params: msg.Msg.GetValue("params"),
	}
	if method, ok := msg.Metadata.GetValue("method").(string); ok {
		rpcp.Method = method
	} else {
		return errors.New("指令方法格式错误")
	}
	var err error
	// 指令下发响应
	if rpcp.Method == "cmdResp" {
		if requestId, ok := msg.Metadata.GetValue("requestId").(string); ok {
			services.DeviceCmdLogModelDao.UpdateResp(requestId, kgo.KConv.ToStr(rpcp.Params), time.Now().Format("2006-01-02 15:04:05"))
		}
	} else {
		result, err := rpcp.GetRequestResult()
		if err != nil {
			if failureLableNode != nil {
				n.Debug(msg, message.DEBUGOUT, err.Error())
				return failureLableNode.Handle(msg)
			} else {
				return err
			}
		}
		// 判断设备协议，根据不通协议，发送不通内容
		deviceProtocol := global.MQTTProtocol
		if msg.Metadata.GetValue("deviceProtocol") != nil && msg.Metadata.GetStringValue("deviceProtocol") != "" {
			deviceProtocol = msg.Metadata.GetValue("deviceProtocol").(string)
		}
		if deviceProtocol == global.MQTTProtocol || deviceProtocol == global.CoAPProtocol || deviceProtocol == global.LwM2MProtocol {
			rpc := &mqttclient.RpcRequest{}
			if n.RequestId == "" {
				if msg.Metadata.GetStringValue("requestId") == "" {
					rpc.RequestId = utils.GenerateID()
				} else {
					rpc.RequestId = msg.Metadata.GetStringValue("requestId")
				}
			} else {
				rpc.RequestId = n.RequestId
			}
			err = rpc.Pub(deviceId, result)
		}
		if deviceProtocol == global.TCPProtocol {
			err = tcpclient.Send(deviceId, result)
		}
		if deviceProtocol == global.UDPProtocol {
			err = udpclient.Send(deviceId, result)
		}
	}

	if err != nil {
		n.Debug(msg, message.DEBUGOUT, err.Error())
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLableNode != nil {
		n.Debug(msg, message.DEBUGOUT, "")
		return successLableNode.Handle(msg)
	}
	return nil
}
