package message

// Message ...
type Message interface {
	GetOriginator() string
	GetType() string
	GetPayload() []byte
	SetType(string)
	SetPayload([]byte)
	SetOriginator(string)
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(b []byte) error
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
		payload: []byte{},
	}
}

type defaultMessage struct {
	originator  string //数据发布者
	messageType string //数据类型，数据来源
	payload     []byte //二进制数据
}

// NewMessageWithDetail ...
func NewMessageWithDetail(originator string, messageType string, payload []byte) Message {
	return &defaultMessage{
		originator:  originator,
		messageType: messageType,
		payload:     payload,
	}
}

func (t *defaultMessage) GetOriginator() string           { return t.originator }
func (t *defaultMessage) GetType() string                 { return t.messageType }
func (t *defaultMessage) GetPayload() []byte              { return t.payload }
func (t *defaultMessage) SetType(messageType string)      { t.messageType = messageType }
func (t *defaultMessage) SetPayload(payload []byte)       { t.payload = payload }
func (t *defaultMessage) SetOriginator(originator string) { t.originator = originator }

func (t *defaultMessage) MarshalBinary() ([]byte, error) { return nil, nil }
func (t *defaultMessage) UnmarshalBinary(b []byte) error { return nil }
