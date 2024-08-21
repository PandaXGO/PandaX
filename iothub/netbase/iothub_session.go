package netbase

import (
	"encoding/json"
	"pandax/pkg/global/model"
)

type DeviceEventInfo struct {
	DeviceId   string            `json:"deviceId"`
	DeviceAuth *model.DeviceAuth `json:"deviceAuth"`
	Datas      string            `json:"datas"`
	Type       string            `json:"type"`
	RequestId  string            `json:"requestId"`  // rpc 请求ID
	Identifier string            `json:"identifier"` //事件标识
}

func (j *DeviceEventInfo) Bytes() []byte {
	b, err := json.Marshal(j)
	if err != nil {
		return nil
	}
	return b
}
