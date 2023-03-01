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
	msgType    string   //消息类型，数据来源
	originator string   //数据发布者
	customerId string   //客户Id  UUID
	entityId   string   //实体Id  UUID
	data       []byte   //数据
	dataType   string   //数据类型   JSON
	metadata   Metadata //数据的元数据
}

// NewMessageWithDetail ...
func NewMessageWithDetail(originator string, messageType string, msg []byte) Message {
	return &defaultMessage{
		originator: originator,
		msgType:    messageType,
		data:       msg,
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

// NewMetadata ...
func NewMetadata() Metadata {
	return &defaultMetadata{
		values: make(map[string]interface{}),
	}
}

type defaultMetadata struct {
	values map[string]interface{}
}

func newDefaultMetadata(vals map[string]interface{}) Metadata {
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
