package netbase

import (
	"pandax/pkg/global_model"
)

type DeviceEventInfo struct {
	DeviceId   string                   `json:"deviceId"`
	DeviceAuth *global_model.DeviceAuth `json:"deviceAuth"`
	Datas      string                   `json:"datas"`
	Type       string                   `json:"type"`
	RequestId  string                   `json:"requestId"`
}
