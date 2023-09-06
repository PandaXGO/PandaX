package message

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// 消息类型
const (
	ConnectMes    = "Connect"
	DisConnectMes = "Disconnect"
	RpcRequestMes = "RpcRequest"
	UpEventMes    = "Event"
	AlarmMes      = "Alarm"
	RowMes        = "Row"
	TelemetryMes  = "Telemetry"
	AttributesMes = "Attributes"
)

// 数据类型Originator
const (
	DEVICE  = "DEVICE"
	GATEWAY = "GATEWAY"
	MONITOR = "MONITOR" //监控
)

type Msg map[string]interface{}
type Metadata map[string]interface{}

type Message struct {
	Id       string    //uuid     消息Id
	Ts       time.Time //时间戳
	MsgType  string    //消息类型，   attributes（参数），telemetry（遥测），Connect连接事件
	UserId   string    //客户Id  UUID  设备发布人
	Msg      Msg       //数据		数据结构JSON  设备原始数据 msg
	Metadata Metadata  //消息的元数据		包括设备Id，设备类型，产品ID等
}

// NewMessage ...
func NewMessage(userId, messageType string, msg Msg, metadata Metadata) *Message {
	return &Message{
		Id:       uuid.New().String(),
		Ts:       time.Now(),
		UserId:   userId,
		MsgType:  messageType,
		Msg:      msg,
		Metadata: metadata,
	}
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
