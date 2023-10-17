package emqxserver

import (
	"context"
	"encoding/json"
	"fmt"
	"pandax/iothub/client/mqttclient"
	"pandax/iothub/hook_message_work"
	"pandax/iothub/netbase"
	exhook2 "pandax/iothub/server/emqxserver/protobuf"
	"pandax/pkg/global"
	"pandax/pkg/global_model"
	"pandax/pkg/rule_engine/message"
	"time"
)

// 创建设备 设备需要一个账号密码 账号使用 namespace.devicename
// 需要创建一个应用事件表? 例如边缘kuiper掉线事件记录或设备事件  type 事件分类

type HookGrpcService struct {
	exhook2.UnimplementedHookProviderServer
	HookService *hook_message_work.HookService
}

func InitEmqxHook(addr string, hs *hook_message_work.HookService) {
	s := NewServer(addr)
	hgs := NewHookGrpcService(hs)
	exhook2.RegisterHookProviderServer(s.GetServe(), hgs)
	err := s.Start(context.TODO())
	if err != nil {
		global.Log.Panic("grpc服务启动错误", err)
	} else {
		global.Log.Infof("IOTHUB HOOK Start SUCCESS,Grpc Server listen: %s", addr)
	}
	// 初始化 MQTT客户端
	mqttclient.InitMqtt(global.Conf.Mqtt.Broker, global.Conf.Mqtt.Username, global.Conf.Mqtt.Password)
}

func NewHookGrpcService(hs *hook_message_work.HookService) *HookGrpcService {
	hgs := &HookGrpcService{
		HookService: hs,
	}
	return hgs
}

func (s *HookGrpcService) OnProviderLoaded(ctx context.Context, in *exhook2.ProviderLoadedRequest) (*exhook2.LoadedResponse, error) {
	hooks := []*exhook2.HookSpec{
		{Name: "client.connect"},
		{Name: "client.connack"},
		{Name: "client.connected"},
		{Name: "client.disconnected"},
		{Name: "client.authenticate"},
		{Name: "client.authorize"},
		{Name: "client.subscribe"},
		{Name: "client.unsubscribe"},
		{Name: "session.created"},
		{Name: "session.subscribed"},
		{Name: "session.unsubscribed"},
		{Name: "session.resumed"},
		{Name: "session.discarded"},
		{Name: "session.takenover"},
		{Name: "session.terminated"},
		{Name: "message.publish"},
		{Name: "message.delivered"},
		{Name: "message.acked"},
		{Name: "message.dropped"},
	}
	return &exhook2.LoadedResponse{Hooks: hooks}, nil
}

func (s *HookGrpcService) OnProviderUnloaded(ctx context.Context, in *exhook2.ProviderUnloadedRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnClientConnect(ctx context.Context, in *exhook2.ClientConnectRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnClientConnack(ctx context.Context, in *exhook2.ClientConnackRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnClientConnected(ctx context.Context, in *exhook2.ClientConnectedRequest) (*exhook2.EmptySuccess, error) {
	global.Log.Info(fmt.Sprintf("Client %s Connected ", in.Clientinfo.GetNode()))
	if in.Clientinfo.Clientid == mqttclient.DefaultDownStreamClientId {
		return &exhook2.EmptySuccess{}, nil
	}
	token := netbase.GetUserName(in.Clientinfo)
	etoken := &global_model.DeviceAuth{}
	etoken.GetDeviceToken(token)
	data := netbase.CreateConnectionInfo(message.ConnectMes, "mqtt", in.Clientinfo.Clientid, in.Clientinfo.Peerhost, etoken)
	s.HookService.MessageCh <- data
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnClientDisconnected(ctx context.Context, in *exhook2.ClientDisconnectedRequest) (*exhook2.EmptySuccess, error) {
	global.Log.Info(fmt.Sprintf("%s断开连接", in.Clientinfo.Username))
	token := netbase.GetUserName(in.Clientinfo)
	if in.Clientinfo.Clientid == mqttclient.DefaultDownStreamClientId {
		return &exhook2.EmptySuccess{}, nil
	}
	etoken := &global_model.DeviceAuth{}
	err := etoken.GetDeviceToken(token)
	if err != nil {
		return nil, err
	}
	data := netbase.CreateConnectionInfo(message.DisConnectMes, "mqtt", in.Clientinfo.Clientid, in.Clientinfo.Peerhost, etoken)
	s.HookService.MessageCh <- data
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnClientAuthenticate(ctx context.Context, in *exhook2.ClientAuthenticateRequest) (*exhook2.ValuedResponse, error) {
	global.Log.Info(fmt.Sprintf("账号%s,开始认证", in.Clientinfo.Username))
	res := &exhook2.ValuedResponse{}
	res.Type = exhook2.ValuedResponse_STOP_AND_RETURN
	res.Value = &exhook2.ValuedResponse_BoolResult{BoolResult: false}

	token := netbase.GetUserName(in.Clientinfo)
	if token == "" {
		global.Log.Warn(fmt.Sprintf("invalid username %s", token))
		return res, nil
	}
	authRes := netbase.Auth(token)
	res.Value = &exhook2.ValuedResponse_BoolResult{BoolResult: authRes}

	return res, nil
}

func (s *HookGrpcService) OnClientAuthorize(ctx context.Context, in *exhook2.ClientAuthorizeRequest) (*exhook2.ValuedResponse, error) {
	reply := &exhook2.ValuedResponse{}
	reply.Type = exhook2.ValuedResponse_STOP_AND_RETURN
	reply.Value = &exhook2.ValuedResponse_BoolResult{BoolResult: true}
	return reply, nil
}

func (s *HookGrpcService) OnClientSubscribe(ctx context.Context, in *exhook2.ClientSubscribeRequest) (*exhook2.EmptySuccess, error) {
	global.Log.Info(fmt.Sprintf("%s订阅了%v", in.Clientinfo.Username, in.TopicFilters))
	// 验证topic 是否是规定的topic，可做topic白名单
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnClientUnsubscribe(ctx context.Context, in *exhook2.ClientUnsubscribeRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnSessionCreated(ctx context.Context, in *exhook2.SessionCreatedRequest) (*exhook2.EmptySuccess, error) {

	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnSessionSubscribed(ctx context.Context, in *exhook2.SessionSubscribedRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnSessionUnsubscribed(ctx context.Context, in *exhook2.SessionUnsubscribedRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnSessionResumed(ctx context.Context, in *exhook2.SessionResumedRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnSessionDiscarded(ctx context.Context, in *exhook2.SessionDiscardedRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnSessionTakenover(ctx context.Context, in *exhook2.SessionTakenoverRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnSessionTerminated(ctx context.Context, in *exhook2.SessionTerminatedRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnMessagePublish(ctx context.Context, in *exhook2.MessagePublishRequest) (*exhook2.ValuedResponse, error) {
	res := &exhook2.ValuedResponse{}
	res.Type = exhook2.ValuedResponse_STOP_AND_RETURN
	res.Value = &exhook2.ValuedResponse_BoolResult{BoolResult: false}

	if in.Message.From == mqttclient.DefaultDownStreamClientId {
		res.Value = &exhook2.ValuedResponse_BoolResult{BoolResult: true}
		return res, nil
	}
	etoken := &global_model.DeviceAuth{}
	etoken.GetDeviceToken(in.Message.Headers["username"])
	// 获取topic类型
	ts := time.Now().Format("2006-01-02 15:04:05.000")
	eventType := IotHubTopic.GetMessageType(in.Message.Topic)

	datas := string(in.GetMessage().GetPayload())
	data := &netbase.DeviceEventInfo{
		Type:       eventType,
		Datas:      datas,
		DeviceId:   etoken.DeviceId,
		DeviceAuth: etoken,
	}

	// 如果是网关子设备单独处理
	if eventType == message.GATEWAY {
		subData := make(map[string]interface{})
		err := json.Unmarshal(in.GetMessage().GetPayload(), &subData)
		if err != nil {
			global.Log.Warn(fmt.Sprintf("子网关上报数据格式错误"))
			res.Value = &exhook2.ValuedResponse_BoolResult{BoolResult: false}
			return res, nil
		}
		// key就是device name
		for deviceName, value := range subData {
			auth, isSub := netbase.SubAuth(deviceName)
			if !isSub {
				continue
			}
			data.DeviceAuth = auth
			data.DeviceId = auth.DeviceId
			if in.Message.Topic == AttributesGatewayTopic {
				data.Type = message.AttributesMes
				marshal, _ := json.Marshal(value)
				attributesData := netbase.UpdateDeviceAttributesData(string(marshal))
				if attributesData == nil {
					continue
				}
				bytes, _ := json.Marshal(attributesData)
				data.Datas = string(bytes)
				// 创建tdengine的设备属性表
				netbase.CreateSubTableField(auth.ProductId, global.TslAttributesType, attributesData)
				// 子设备发送到队列里
				s.HookService.MessageCh <- data
			}
			if in.Message.Topic == TelemetryGatewayTopic {
				data.Type = message.TelemetryMes
				// 数据处理 如果上传的数据没有时间戳 添加时间戳更改格式化
				td, _ := json.Marshal(value)
				telemetryData := netbase.UpdateDeviceTelemetryData(string(td))
				if telemetryData == nil {
					continue
				}
				bytes, _ := json.Marshal(telemetryData)
				data.Datas = string(bytes)
				// 创建tdengine的设备遥测表
				netbase.CreateSubTableField(auth.ProductId, global.TslTelemetryType, telemetryData)
				// 子设备发送到队列里
				s.HookService.MessageCh <- data
			}
			if in.Message.Topic == ConnectGatewayTopic {
				if val, ok := value.(string); ok {
					if val == "online" {
						data = netbase.CreateConnectionInfo(message.ConnectMes, "mqtt", in.Message.From, in.Message.Headers["peerhost"], auth)
					}
					if val == "offline" {
						data = netbase.CreateConnectionInfo(message.DisConnectMes, "mqtt", in.Message.From, in.Message.Headers["peerhost"], auth)
					}
					// 子设备发送到队列里
					s.HookService.MessageCh <- data
				}
			}
		}
		res.Value = &exhook2.ValuedResponse_Message{Message: in.Message}
		return res, nil
	}

	switch eventType {
	case message.RowMes:
		data.Type = message.RowMes
		data.Datas = fmt.Sprintf(`{"ts": "%s","rowdata": "%s"}`, ts, data.Datas)
	case message.AttributesMes:
		attributesData := netbase.UpdateDeviceAttributesData(data.Datas)
		if attributesData == nil {
			return res, nil
		}
		bytes, _ := json.Marshal(attributesData)
		data.Datas = string(bytes)
	case message.TelemetryMes:
		// 数据处理 如果上传的数据没有时间戳 添加时间戳更改格式化
		telemetryData := netbase.UpdateDeviceTelemetryData(data.Datas)
		if telemetryData == nil {
			return res, nil
		}
		bytes, _ := json.Marshal(telemetryData)
		data.Datas = string(bytes)
	case message.RpcRequestFromDevice:
		// 获取请求id
		id := netbase.GetRequestIdFromTopic(RpcReq, in.Message.Topic)
		data.RequestId = id
	}

	//TODO 如果设备消息；量过大，推荐采用NATS队列处理
	s.HookService.MessageCh <- data

	res.Value = &exhook2.ValuedResponse_Message{Message: in.Message}
	return res, nil
}

func (s *HookGrpcService) OnMessageDelivered(ctx context.Context, in *exhook2.MessageDeliveredRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnMessageDropped(ctx context.Context, in *exhook2.MessageDroppedRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}

func (s *HookGrpcService) OnMessageAcked(ctx context.Context, in *exhook2.MessageAckedRequest) (*exhook2.EmptySuccess, error) {
	return &exhook2.EmptySuccess{}, nil
}
