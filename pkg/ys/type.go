package ys

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  int64  `json:"expireTime"`
}

type Status struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Page interface{} `json:"page"`
}

type Page struct {
	Total float64 `json:"total"`
	Page  float64 `json:"page"`
	Size  float64 `json:"size"`
}

// Device 萤石设备数据结构
type Device struct {
	DeviceSerial  string `json:"deviceSerial"`
	DeviceName    string `json:"deviceName"`
	DeviceType    string `json:"deviceType"`
	Status        int    `json:"status"`
	Defence       int    `json:"defence"`
	DeviceVersion string `json:"deviceVersion"`
	NetAddress    string `json:"netAddress"`
}

// Channel 萤石摄像头通道数据结构
type Channel struct {
	DeviceSerial string `json:"deviceSerial"`
	IpcSerial    string `json:"ipcSerial"`
	ChannelNo    int    `json:"channelNo"`
	ChannelName  string `json:"channelName"`
	PicURL       string `json:"picUrl"`
	IsShared     string `json:"isShared"`
	VideoLevel   int    `json:"videoLevel"`
	IsEncrypt    int    `json:"isEncrypt"`
	Status       int    `json:"status"`
}

// LiveAddress 播放地址
type LiveAddress struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	ExpireTime string `json:"expireTime"`
}
