package message

import "github.com/sirupsen/logrus"

// Message ...
type Message interface {
	GetOriginator() string
	GetType() string
	GetData() []byte
	GetMetadata() Metadata
	SetType(string)
	SetData([]byte)
	SetOriginator(string)
	SetMetadata(Metadata)
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(b []byte) error
}

// Metadata ...
type Metadata interface {
	Keys() []string
	GetKeyValue(key string) interface{}
	SetKeyValue(key string, val interface{})
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
		data: []byte{},
	}
}

type defaultMessage struct {
	id         string   //uuid
	ts         int64    //时间戳
	msgType    string   //消息类型，   attributes（参数），telemetry（遥测），连接事件
	originator string   //数据发布者
	customerId string   //客户Id  UUID
	entityId   string   //实体Id  UUID
	entityType string   //实体类型  设备、资产，用户、规则链
	data       []byte   //数据		数据结构JSON   设备原始数据
	metadata   Metadata //消息的元数据  包括，设备名称，设备类型，命名空间，时间戳等
}

// NewMessageWithDetail ...
func NewMessageWithDetail(originator string, messageType string, msg []byte, metadata Metadata) Message {
	return &defaultMessage{
		originator: originator,
		msgType:    messageType,
		data:       msg,
		metadata:   metadata,
	}
}

func (t *defaultMessage) GetOriginator() string           { return t.originator }
func (t *defaultMessage) GetType() string                 { return t.msgType }
func (t *defaultMessage) GetData() []byte                 { return t.data }
func (t *defaultMessage) GetMetadata() Metadata           { return t.metadata }
func (t *defaultMessage) SetType(msgType string)          { t.msgType = msgType }
func (t *defaultMessage) SetData(data []byte)             { t.data = data }
func (t *defaultMessage) SetOriginator(originator string) { t.originator = originator }
func (t *defaultMessage) SetMetadata(metadata Metadata)   { t.metadata = metadata }

func (t *defaultMessage) MarshalBinary() ([]byte, error) { return nil, nil }
func (t *defaultMessage) UnmarshalBinary(b []byte) error { return nil }

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
