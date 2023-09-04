package iothub

import (
	"context"
	"encoding/json"
	"fmt"
	exhook "pandax/iothub/protobuf"
	"pandax/pkg/global"
	"pandax/pkg/mqtt"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/tdengine"
	"pandax/pkg/tool"
	"strconv"
	"sync"
	"time"
)

// 创建设备 设备需要一个账号密码 账号使用 namespace.devicename
// 需要创建一个应用事件表? 例如边缘kuiper掉线事件记录或设备事件  type 事件分类

type HookService struct {
	exhook.UnimplementedHookProviderServer
	cache     sync.Map
	wg        sync.WaitGroup //  优雅关闭
	ch        chan struct{}  //  并发限制
	messageCh chan *DeviceEventInfo
}

func InitEmqxHook(addr string) *HookService {
	s := NewServer(addr)
	service := NewHookService()
	exhook.RegisterHookProviderServer(s.GetServe(), service)
	err := s.Start(context.TODO())
	if err != nil {
		global.Log.Panic("grpc服务启动错误", err)
	} else {
		global.Log.Infof("IOTHUB HOOK Start SUCCESS,Grpc Server listen: %s", addr)
	}
	// 开启线程处理消息
	go service.MessageWork()
	// 初始化 MQTT客户端
	global.MqttClient = mqtt.InitMqtt(global.Conf.Mqtt.Broker, global.Conf.Mqtt.Username, global.Conf.Mqtt.Password)
	return service
}

func NewHookService() *HookService {
	hs := &HookService{
		cache:     sync.Map{},
		messageCh: make(chan *DeviceEventInfo),
	}
	// 并发限制，代表服务器处理能力
	if global.Conf.Queue.Enable && global.Conf.Queue.Num > 0 {
		hs.ch = make(chan struct{}, global.Conf.Queue.Num)
	}
	return hs
}

func (s *HookService) OnProviderLoaded(ctx context.Context, in *exhook.ProviderLoadedRequest) (*exhook.LoadedResponse, error) {
	hooks := []*exhook.HookSpec{
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
	return &exhook.LoadedResponse{Hooks: hooks}, nil
}

func (s *HookService) OnProviderUnloaded(ctx context.Context, in *exhook.ProviderUnloadedRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnClientConnect(ctx context.Context, in *exhook.ClientConnectRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnClientConnack(ctx context.Context, in *exhook.ClientConnackRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnClientConnected(ctx context.Context, in *exhook.ClientConnectedRequest) (*exhook.EmptySuccess, error) {
	global.Log.Info(fmt.Sprintf("Client %s Connected ", in.Clientinfo.GetNode()))
	ts := time.Now().Format("2006-01-02 15:04:05.000")
	username := GetUserName(in.Clientinfo)
	ci := &tdengine.ConnectInfo{
		ClientID:   in.Clientinfo.Clientid,
		DeviceId:   username,
		PeerHost:   in.Clientinfo.Peerhost,
		Protocol:   in.Clientinfo.Protocol,
		SocketPort: strconv.Itoa(int(in.Clientinfo.Sockport)),
		Type:       message.ConnectMes,
		Ts:         ts,
	}
	v, err := EncodeData(*ci)
	if err != nil {
		return nil, err
	}
	// 添加设备上线记录
	data := &DeviceEventInfo{
		DeviceId: username,
		Datas:    string(v),
		Type:     message.ConnectMes,
	}
	s.messageCh <- data

	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnClientDisconnected(ctx context.Context, in *exhook.ClientDisconnectedRequest) (*exhook.EmptySuccess, error) {
	global.Log.Info(fmt.Sprintf("%s断开连接", in.Clientinfo.Username))
	devicename := GetUserName(in.Clientinfo)

	ts := time.Now().Format("2006-01-02 15:04:05.000")
	ci := &tdengine.ConnectInfo{
		ClientID:   in.Clientinfo.Clientid,
		DeviceId:   devicename,
		PeerHost:   in.Clientinfo.Peerhost,
		Protocol:   in.Clientinfo.Protocol,
		SocketPort: strconv.Itoa(int(in.Clientinfo.Sockport)),
		Type:       message.DisConnectMes,
		Ts:         ts,
	}
	v, err := EncodeData(*ci)
	if err != nil {
		return nil, err
	}

	// 添加设备下线记录
	data := &DeviceEventInfo{
		DeviceId: devicename,
		Datas:    string(v),
		Type:     message.DisConnectMes,
	}
	s.messageCh <- data
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnClientAuthenticate(ctx context.Context, in *exhook.ClientAuthenticateRequest) (*exhook.ValuedResponse, error) {
	global.Log.Info(fmt.Sprintf("账号%s，密码%s,开始认证", in.Clientinfo.Username, in.Clientinfo.Password))
	res := &exhook.ValuedResponse{}
	res.Type = exhook.ValuedResponse_STOP_AND_RETURN
	res.Value = &exhook.ValuedResponse_BoolResult{BoolResult: false}

	username := GetUserName(in.Clientinfo)
	pw := GetPassword(in.Clientinfo)
	if username == "" || pw == "" {
		global.Log.Warn(fmt.Sprintf("invalid username %s or password %s", username, pw))
		return res, nil
	}
	authRes := s.auth(username, pw)
	res.Value = &exhook.ValuedResponse_BoolResult{BoolResult: authRes}

	return res, nil
}

func (s *HookService) OnClientAuthorize(ctx context.Context, in *exhook.ClientAuthorizeRequest) (*exhook.ValuedResponse, error) {
	reply := &exhook.ValuedResponse{}
	reply.Type = exhook.ValuedResponse_STOP_AND_RETURN
	reply.Value = &exhook.ValuedResponse_BoolResult{BoolResult: true}
	return reply, nil
}

func (s *HookService) OnClientSubscribe(ctx context.Context, in *exhook.ClientSubscribeRequest) (*exhook.EmptySuccess, error) {
	global.Log.Info(fmt.Sprintf("%s订阅了%v", in.Clientinfo.Username, in.TopicFilters))
	// 验证topic 是否是规定的topic，可做topic白名单
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnClientUnsubscribe(ctx context.Context, in *exhook.ClientUnsubscribeRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnSessionCreated(ctx context.Context, in *exhook.SessionCreatedRequest) (*exhook.EmptySuccess, error) {

	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnSessionSubscribed(ctx context.Context, in *exhook.SessionSubscribedRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnSessionUnsubscribed(ctx context.Context, in *exhook.SessionUnsubscribedRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnSessionResumed(ctx context.Context, in *exhook.SessionResumedRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnSessionDiscarded(ctx context.Context, in *exhook.SessionDiscardedRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnSessionTakenover(ctx context.Context, in *exhook.SessionTakenoverRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnSessionTerminated(ctx context.Context, in *exhook.SessionTerminatedRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnMessagePublish(ctx context.Context, in *exhook.MessagePublishRequest) (*exhook.ValuedResponse, error) {
	res := &exhook.ValuedResponse{}
	res.Type = exhook.ValuedResponse_STOP_AND_RETURN
	res.Value = &exhook.ValuedResponse_BoolResult{BoolResult: false}

	if in.Message.From == mqtt.DefaultDownStreamClientId {
		res.Value = &exhook.ValuedResponse_BoolResult{BoolResult: true}
		return res, nil
	}
	// 获取topic类型
	ts := time.Now().Format("2006-01-02 15:04:05.000")
	eventType := IotHubTopic.GetMessageType(in.Message.Topic)
	datas := string(in.GetMessage().GetPayload())
	data := &DeviceEventInfo{
		Type:     eventType,
		Datas:    datas,
		DeviceId: in.Message.Headers["username"],
	}
	// 如果是网关子设备单独处理
	if eventType == message.GATEWAY {
		subData := make(map[string]interface{})
		err := json.Unmarshal(in.GetMessage().GetPayload(), &subData)
		if err != nil {
			global.Log.Warn(fmt.Sprintf("子网关上报数据格式错误"))
			res.Value = &exhook.ValuedResponse_BoolResult{BoolResult: false}
			return res, nil
		}
		// key就是deviceId
		for key, value := range subData {
			etoken := &tool.DeviceAuth{}
			err = global.RedisDb.Get(key, etoken)
			if err != nil {
				global.Log.Infof("%s设备不存在", key)
				continue
			}
			data.DeviceId = key
			if in.Message.Topic == AttributesGatewayTopic {
				data.Type = message.AttributesMes
				marshal, _ := json.Marshal(value)
				attributesData := updateDeviceAttributesData(string(marshal))
				if attributesData == nil {
					continue
				}
				bytes, _ := json.Marshal(attributesData)
				data.Datas = string(bytes)
				// 子设备发送到队列里
				s.messageCh <- data
			}
			if in.Message.Topic == TelemetryGatewayTopic {
				data.Type = message.TelemetryMes
				// 数据处理 如果上传的数据没有时间戳 添加时间戳更改格式化
				telData := make([]map[string]interface{}, 0)
				// 解析子设备遥测数据结构
				marshal, _ := json.Marshal(value)
				err := json.Unmarshal(marshal, &telData)
				if err != nil {
					global.Log.Infof("%s子设备遥测数据结构错误", key)
					continue
				}
				for _, da := range telData {
					td, _ := json.Marshal(da)
					telemetryData := updateDeviceTelemetryData(string(td))
					if telemetryData == nil {
						continue
					}
					bytes, _ := json.Marshal(telemetryData)
					data.Datas = string(bytes)
					// 子设备发送到队列里
					s.messageCh <- data
				}
			}
			if in.Message.Topic == ConnectGatewayTopic {
				data.Type = message.ConnectMes
				ci := &tdengine.ConnectInfo{
					ClientID: in.Message.From,
					Protocol: in.Message.Headers["protocol"],
					PeerHost: in.Message.Headers["peerhost"],
					DeviceId: key,
					Type:     message.ConnectMes,
					Ts:       ts,
				}
				v, _ := EncodeData(*ci)
				data.Datas = string(v)
				// 子设备发送到队列里
				s.messageCh <- data
			}
			if in.Message.Topic == DisconnectGatewayTopic {
				data.Type = message.DisConnectMes
				ci := &tdengine.ConnectInfo{
					ClientID: in.Message.From,
					DeviceId: key,
					Protocol: in.Message.Headers["protocol"],
					PeerHost: in.Message.Headers["peerhost"],
					Type:     message.DisConnectMes,
					Ts:       ts,
				}
				v, _ := EncodeData(*ci)
				data.Datas = string(v)
				// 子设备发送到队列里
				s.messageCh <- data
			}
		}
		res.Value = &exhook.ValuedResponse_Message{Message: in.Message}
		return res, nil
	}

	switch eventType {
	case message.RowMes:
		data.Datas = fmt.Sprintf(`{"ts": "%s","rowdata": "%s"}`, ts, data.Datas)
	case message.AttributesMes:
		attributesData := updateDeviceAttributesData(data.Datas)
		if attributesData == nil {
			return res, nil
		}
		bytes, _ := json.Marshal(attributesData)
		data.Datas = string(bytes)
	case message.TelemetryMes:
		// 数据处理 如果上传的数据没有时间戳 添加时间戳更改格式化
		telemetryData := updateDeviceTelemetryData(data.Datas)
		if telemetryData == nil {
			return res, nil
		}
		bytes, _ := json.Marshal(telemetryData)
		data.Datas = string(bytes)
	case message.RpcRequestMes:
		// 获取请求id
		id := getRequestIdFromTopic(RpcReqReg, in.Message.Topic)
		data.RequestId = id
	}

	//TODO 如果设备消息；量过大，推荐采用NATS队列处理
	s.messageCh <- data

	res.Value = &exhook.ValuedResponse_Message{Message: in.Message}
	return res, nil
}

func (s *HookService) OnMessageDelivered(ctx context.Context, in *exhook.MessageDeliveredRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnMessageDropped(ctx context.Context, in *exhook.MessageDroppedRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}

func (s *HookService) OnMessageAcked(ctx context.Context, in *exhook.MessageAckedRequest) (*exhook.EmptySuccess, error) {
	return &exhook.EmptySuccess{}, nil
}
func (s *HookService) Stop() {
	s.wg.Wait()
}
