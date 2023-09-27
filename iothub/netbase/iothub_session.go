package netbase

import "pandax/pkg/tool"

type DeviceEventInfo struct {
	DeviceId   string           `json:"deviceId"`
	DeviceAuth *tool.DeviceAuth `json:"deviceAuth"`
	Datas      string           `json:"datas"`
	Type       string           `json:"type"`
	RequestId  string           `json:"requestId"`
}
