package ys

const (
	//云台控制开始
	URLPTZSTAR = "https://open.ys7.com/api/lapp/device/ptz/start"
	//云台控制结束
	URLPTZSTOP = "https://open.ys7.com/api/lapp/device/ptz/stop"
)

// StartPtz 开始云台控制
func (ys *Ys) StartPtz(deviceSerial string, channelNo, direction, speed int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["direction"] = direction
	params["speed"] = speed

	_, err = ys.authorizeRequset("POST", URLPTZSTAR, params, nil)
	if err != nil {
		return
	}
	return nil
}

// StopPtz 停止云台转动
func (ys *Ys) StopPtz(deviceSerial string, channelNo, direction int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["direction"] = direction

	_, err = ys.authorizeRequset("POST", URLPTZSTOP, params, nil)
	if err != nil {
		return
	}
	return nil
}
