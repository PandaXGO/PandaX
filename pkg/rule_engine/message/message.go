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

// Message ...
type Message interface {
	GetId() string
	GetTs() time.Time
	GetUserId() string
	GetType() string
	GetMsg() map[string]interface{}
	GetMetadata() Metadata
	GetAllMap() map[string]interface{} //msg 和 Metadata的合并
	SetType(string)
	SetMsg(map[string]interface{})
	SetUserId(string)
	SetMetadata(Metadata)
	MarshalBinary() ([]byte, error)
}

// Metadata ...
type Metadata interface {
	Keys() []string
	GetKeyValue(key string) interface{}
	SetKeyValue(key string, val interface{})
	GetValues() map[string]interface{}
}

// NewMessage ...
func NewMessage() Message {
	return &defaultMessage{
		Id:  uuid.New().String(),
		Ts:  time.Now(),
		Msg: map[string]interface{}{},
	}
}

type defaultMessage struct {
	Id       string                 //uuid     消息Id
	Ts       time.Time              //时间戳
	MsgType  string                 //消息类型，   attributes（参数），telemetry（遥测），Connect连接事件
	UserId   string                 //客户Id  UUID  设备发布人
	Msg      map[string]interface{} //数据		数据结构JSON  设备原始数据 msg
	Metadata Metadata               //消息的元数据		包括设备Id，设备类型，产品ID等
}

// NewMessageWithDetail ...
func NewMessageWithDetail(userId, messageType string, msg map[string]interface{}, metadata Metadata) Message {
	return &defaultMessage{
		Id:       uuid.New().String(),
		Ts:       time.Now(),
		UserId:   userId,
		MsgType:  messageType,
		Msg:      msg,
		Metadata: metadata,
	}
}

func (t *defaultMessage) GetId() string                  { return t.Id }
func (t *defaultMessage) GetTs() time.Time               { return t.Ts }
func (t *defaultMessage) GetUserId() string              { return t.UserId }
func (t *defaultMessage) GetType() string                { return t.MsgType }
func (t *defaultMessage) GetMsg() map[string]interface{} { return t.Msg }
func (t *defaultMessage) GetMetadata() Metadata          { return t.Metadata }
func (t *defaultMessage) GetAllMap() map[string]interface{} {
	data := make(map[string]interface{})
	for msgKey, msgValue := range t.GetMsg() {
		for metaKey, metaValue := range t.GetMetadata().GetValues() {
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
func (t *defaultMessage) SetType(msgType string)            { t.MsgType = msgType }
func (t *defaultMessage) SetMsg(msg map[string]interface{}) { t.Msg = msg }
func (t *defaultMessage) SetUserId(userId string)           { t.UserId = userId }
func (t *defaultMessage) SetMetadata(metadata Metadata)     { t.Metadata = metadata }

func (t *defaultMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

// NewMetadata ...
func NewMetadata() Metadata {
	return &defaultMetadata{
		values: make(map[string]interface{}),
	}
}

type defaultMetadata struct {
	values map[string]interface{}
}

func NewDefaultMetadata(vals map[string]interface{}) Metadata {
	return &defaultMetadata{
		values: vals,
	}
}

func (t *defaultMetadata) Keys() []string {
	keys := make([]string, 0)
	for key := range t.values {
		keys = append(keys, key)
	}
	return keys
}

func (t *defaultMetadata) GetKeyValue(key string) interface{} {
	if _, found := t.values[key]; !found {
		return nil
	}
	return t.values[key]
}

func (t *defaultMetadata) SetKeyValue(key string, val interface{}) {
	t.values[key] = val
}

func (t *defaultMetadata) GetValues() map[string]interface{} {
	return t.values
}
func (t *defaultMetadata) SetValues(values map[string]interface{}) {
	t.values = values
}
