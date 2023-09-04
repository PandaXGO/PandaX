package ys

import (
	"pandax/pkg/tool"
)

const (
	//设备列表
	DEVICELIST = "https://open.ys7.com/api/lapp/device/list"
	//获取指定设备的通道信息
	DEVICECHANNELLIST = "https://open.ys7.com/api/lapp/device/camera/list"
	// 获取播放地址
	DEVICELIVEADDRESS = "https://open.ys7.com/api/lapp/v2/live/address/get"
)

// GetDeviceList 获取设备列表
func (ys *Ys) GetDeviceList(pageNum, pageSize int) (devices []Device, total int64, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageNum
	params["pageSize"] = pageSize
	status, err := ys.authorizeRequset("POST", DEVICELIST, params, &devices) //获取用户下的设备列表
	if err != nil {
		return nil, 0, err
	}
	var page Page
	err = tool.InterfaceToStruct(status.Page, &page)
	if err != nil {
		return nil, 0, err
	}
	return devices, int64(page.Total), nil
}

// GetDeviceChannelList 获取指定设备的通道信息
func (ys *Ys) GetDeviceChannelList(deviceSerial string) (cameras []Channel, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", DEVICECHANNELLIST, params, &cameras)
	if err != nil {
		return nil, err
	}
	return cameras, nil
}

// GetDeviceLiveAddress 获取指定设备通道的播放地址
func (ys *Ys) GetDeviceLiveAddress(deviceSerial string, channelNo int) (live []LiveAddress, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["protocol"] = 1 //流播放协议，1-ezopen、2-hls、3-rtmp、4-flv，默认为1
	params["type"] = "1"   //地址的类型，1-预览，2-本地录像回放，3-云存储录像回放，非必选，默认为1；回放仅支持rtmp、ezopen、flv协议
	params["quality"] = 1  //视频清晰度，1-高清（主码流）、2-流畅（子码流）
	_, err = ys.authorizeRequset("POST", DEVICELIVEADDRESS, params, &live)
	if err != nil {
		return nil, err
	}
	return live, nil
}
