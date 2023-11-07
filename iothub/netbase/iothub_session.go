package netbase

import (
	"pandax/pkg/global/model"
)

type DeviceEventInfo struct {
	DeviceId   string            `json:"deviceId"`
	DeviceAuth *model.DeviceAuth `json:"deviceAuth"`
	Datas      string            `json:"datas"`
	Type       string            `json:"type"`
	RequestId  string            `json:"requestId"`
}
