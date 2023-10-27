package message

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

// 消息类型
const (
	ConnectMes           = "Connect"
	DisConnectMes        = "Disconnect"
	RpcRequestFromDevice = "RpcRequestFromDevice"
	RpcRequestToDevice   = "RpcRequestToDevice"
	UpEventMes           = "Event"
	AlarmMes             = "Alarm"
	RowMes               = "Row"
	TelemetryMes         = "Telemetry"
	AttributesMes        = "Attributes"
)

// 数据类型Originator
const (
	DEVICE  = "DEVICE"
	GATEWAY = "GATEWAY"
	MONITOR = "MONITOR" //监控
)

const (
	DEBUGIN  = "In"
	DEBUGOUT = "Out"
)

type Msg map[string]interface{}
type Metadata map[string]interface{}

type Message struct {
	Id           string    //uuid     消息Id
	Ts           time.Time //时间戳
	MsgType      string    //消息类型，   attributes（参数），telemetry（遥测），Connect连接事件
	User         string    //客户  设备发布人 设备所有者
	Msg          Msg       //数据		数据结构JSON  设备原始数据 msg
	Metadata     Metadata  //消息的元数据		包括设备Id，设备类型，产品ID等
	DeBugChan    chan DebugData
	EndDeBugChan chan struct{}
}

// NewMessage ...
func NewMessage(user, messageType string, msg Msg, metadata Metadata) *Message {
	return &Message{
		Id:           uuid.New().String(),
		Ts:           time.Now(),
		User:         user,
		MsgType:      messageType,
		Msg:          msg,
		Metadata:     metadata,
		DeBugChan:    make(chan DebugData, 100),
		EndDeBugChan: make(chan struct{}),
	}
}

func (t *Message) Debug(nodeId, nodeName, debugType, error string) {
	if debugType == DEBUGIN {
		logrus.Infof("%s handle message '%s'", nodeName, t.MsgType)
	}
	debug := DebugData{
		Ts:        time.Now().Format("2006-01-02 15:04:05.000"),
		NodeId:    nodeId,
		MsgId:     t.Id,
		DebugType: debugType,
		MsgType:   t.MsgType,
		Msg:       t.Msg,
		Metadata:  t.Metadata,
		Error:     error,
	}
	if deviceName, ok := t.Metadata.GetValue("deviceName").(string); ok {
		debug.DeviceName = deviceName
	}
	t.DeBugChan <- debug
}

func (t *Message) GetAllMap() map[string]interface{} {
	data := make(map[string]interface{})
	for msgKey, msgValue := range t.Msg {
		for metaKey, metaValue := range t.Metadata {
			if msgKey == metaKey {
				data[msgKey] = metaValue
			} else {
				if _, ok := data[msgKey]; !ok {
					data[msgKey] = msgValue
				}
				if _, ok := data[metaKey]; !ok {
					data[metaKey] = metaValue
				}
			}
		}
	}
	return data
}

func (t *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (meta *Metadata) Keys() []string {
	keys := make([]string, 0)
	for key := range *meta {
		keys = append(keys, key)
	}
	return keys
}

func (meta *Metadata) GetValue(key string) any {
	if _, found := (*meta)[key]; !found {
		return nil
	}
	return (*meta)[key]
}

func (meta *Metadata) SetValue(key string, val interface{}) {
	(*meta)[key] = val
}

func (msg *Msg) GetValue(key string) any {
	if _, found := (*msg)[key]; !found {
		return nil
	}
	return (*msg)[key]
}
