package iothub

// 指令下发
/*func Control(assets, thingModel, device_name, parameter string, operation bool) error {
	topic := fmt.Sprintf("control/%s/%s", assets, device_name)
	log.Println(topic)
	payload := fmt.Sprintf(`{"method":"control","data":{"parameter": "%s","operation":%t}}`, parameter, operation)
	//Publish(*global.GVA_MQTT, topic, 1, payload)
	return nil
}

func ControlState(assets, thingModel, device_name string) (map[string]interface{}, error) {
	topic := fmt.Sprintf("control/%s/%s", assets, device_name)
	payload := fmt.Sprintf(`{"method":"state","data":{}}`)
	if Publish(*global.GVA_MQTT, topic, 1, payload) != nil {
		return nil, errors.New("下发获取状态参数指令失败")
	}
	select {
	case state := <-controlState:
		return state, nil
	case <-time.After(10 * time.Second):
		return nil, errors.New("请求指令状态超时")
	}
}
*/
