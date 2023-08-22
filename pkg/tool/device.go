package tool

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"pandax/pkg/global"
	"strconv"
	"strings"
)

type DeviceAuth struct {
	User        string `json:"user"`
	DeviceId    string `json:"device_id"`
	DeviceType  string `json:"device_type"`
	ProductId   string `json:"product_id"`
	RuleChainId string `json:"rule_chain_id"`
	Name        string `json:"name"`
	Token       string `json:"token"`
	CreatedAt   int64  `json:"created_at"`
	ExpiredAt   int64  `json:"expired_at"`
}

func (entity *DeviceAuth) CreateDeviceToken() (err error) {

	return nil
}

func (entity *DeviceAuth) GetDeviceToken(key string) error {
	if err := global.RedisDb.Get(key, entity); err != nil {
		return err
	}
	return nil
}

func (token *DeviceAuth) MD5ID() string {
	buf := bytes.NewBufferString(token.DeviceId)
	buf.WriteString(token.DeviceType)
	buf.WriteString(strconv.FormatInt(token.CreatedAt, 10))
	access := base64.URLEncoding.EncodeToString([]byte(uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes()).String()))
	access = strings.TrimRight(access, "=")
	return access
}
func (token *DeviceAuth) GetMarshal() string {
	marshal, _ := json.Marshal(*token)
	return string(marshal)
}

func (token *DeviceAuth) GetUnMarshal(data []byte) error {
	return json.Unmarshal(data, token)
}

// 序列化
func (m *DeviceAuth) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// 反序列化
func (m *DeviceAuth) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
