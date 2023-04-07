package message

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

// Message ...
type Message interface {
	GetOriginator() string
	GetType() string
	GetMsg() map[string]interface{}
	GetMetadata() Metadata
	SetType(string)
	SetMsg(map[string]interface{})
	SetOriginator(string)
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

// Predefined message types
const (
	MessageTypePostAttributesRequest = "Post attributes"
	MessageTypePostTelemetryRequest  = "Post telemetry"
	MessageTypeActivityEvent         = "Activity event"
	MessageTypeInactivityEvent       = "Inactivity event"
	MessageTypeConnectEvent          = "Connect event"
	MessageTypeDisconnectEvent       = "Disconnect event"
)

// NewMessage ...
func NewMessage() Message {
	return &defaultMessage{
		id:  uuid.New().String(),
		ts:  time.Now(),
		msg: map[string]interface{}{},
	}
}

type defaultMessage struct {
	id         string                 //uuid
	ts         time.Time              //时间戳
	msgType    string                 //消息类型，   attributes（参数），telemetry（遥测），Connect连接事件
	originator string                 //数据发布者	 设备 规则链
	userId     string                 //客户Id  UUID
	deviceId   string                 //设备Id  UUID
	msg        map[string]interface{} //数据		数据结构JSON  设备原始数据 msg
	metadata   Metadata               //消息的元数据		包括，设备名称，设备类型，命名空间，时间戳等
}

// NewMessageWithDetail ...
func NewMessageWithDetail(originator string, messageType string, msg map[string]interface{}, metadata Metadata) Message {
	return &defaultMessage{
		originator: originator,
		msgType:    messageType,
		msg:        msg,
		metadata:   metadata,
	}
}

func (t *defaultMessage) GetOriginator() string             { return t.originator }
func (t *defaultMessage) GetType() string                   { return t.msgType }
func (t *defaultMessage) GetMsg() map[string]interface{}    { return t.msg }
func (t *defaultMessage) GetMetadata() Metadata             { return t.metadata }
func (t *defaultMessage) SetType(msgType string)            { t.msgType = msgType }
func (t *defaultMessage) SetMsg(msg map[string]interface{}) { t.msg = msg }
func (t *defaultMessage) SetOriginator(originator string)   { t.originator = originator }
func (t *defaultMessage) SetMetadata(metadata Metadata)     { t.metadata = metadata }
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
		logrus.Fatalf("no key '%s' in metadata", key)
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
