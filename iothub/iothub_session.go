package iothub

type DeviceEventInfo struct {
	DeviceId  string `json:"deviceId"`
	Datas     string `json:"datas"`
	Type      string `json:"type"`
	RequestId string `json:"requestId"`
}
