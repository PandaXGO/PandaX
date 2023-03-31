package message

import "github.com/sirupsen/logrus"

// Message ...
type Message interface {
	GetOriginator() string
	GetType() string
	GetCh() chan map[string]bool
	GetMetadata() Metadata
	SetType(string)
	SetOriginator(string)
	SetMetadata(Metadata)
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
		ch: make(chan map[string]bool),
	}
}

type defaultMessage struct {
	id         string               //uuid
	ts         int64                //时间戳
	msgType    string               //消息类型
	originator string               //数据发布者
	ch         chan map[string]bool //通道
	metadata   Metadata             //消息的元数据		包括，设备名称，设备类型，命名空间，时间戳等
}

// NewMessageWithDetail ...
func NewMessageWithDetail(originator string, messageType string, msg map[string]interface{}, metadata Metadata) Message {
	return &defaultMessage{
		originator: originator,
		msgType:    messageType,
		ch:         make(chan map[string]bool),
		metadata:   metadata,
	}
}

func (t *defaultMessage) GetOriginator() string           { return t.originator }
func (t *defaultMessage) GetType() string                 { return t.msgType }
func (t *defaultMessage) GetCh() chan map[string]bool     { return t.ch }
func (t *defaultMessage) GetMetadata() Metadata           { return t.metadata }
func (t *defaultMessage) SetType(msgType string)          { t.msgType = msgType }
func (t *defaultMessage) SetOriginator(originator string) { t.originator = originator }
func (t *defaultMessage) SetMetadata(metadata Metadata)   { t.metadata = metadata }

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
