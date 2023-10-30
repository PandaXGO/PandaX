package mqttclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"pandax/pkg/global"
)

// key 设备id，value MQTT的clientID
var MqttClient = make(map[string]string)

const ClientsInfo string = "client"
const SubscribeTopicsInfo string = "subscribe"

// GetEmqInfo 获取emqx信息，包括连接信息，订阅信息
func GetEmqInfo(infoType string) ([]map[string]interface{}, error) {
	var url string
	if infoType == ClientsInfo {
		url = fmt.Sprintf("%s/v5/clients?_page=1&_limit=100000", global.Conf.Mqtt.Broker)
	} else if infoType == SubscribeTopicsInfo {
		url = fmt.Sprintf("%s/v5/subscriptions?_page=1&_limit=100000", global.Conf.Mqtt.Broker)
	} else {
		return nil, errors.New("invalid infoType")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if nil != err {
		return nil, err
	}
	req.SetBasicAuth(global.Conf.Mqtt.Username, global.Conf.Mqtt.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	global.Log.Debug("receive resp, ", string(body))
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var result interface{}
	if err = json.Unmarshal(body, &result); nil != err {
		global.Log.Error("body Unmarshal error", err)
		return nil, err
	}
	res, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New("result error")
	}
	data := res["data"].([]map[string]interface{})
	return data, nil
}

// Publish 推送信息
func Publish(topic, clientId string, payload interface{}) error {
	if clientId == "" {
		return errors.New("未获取到MQTT连接")
	}
	global.Log.Debugf("send data to clientId: %s, topic:%s, payload: %v", clientId, topic, payload)
	url := fmt.Sprintf("%s/v5/publish", global.Conf.Mqtt.Broker)
	pubData := map[string]interface{}{
		"topic":    topic,
		"payload":  payload,
		"qos":      global.Conf.Mqtt.Qos,
		"retain":   false,
		"clientid": clientId,
	}
	data, err := json.Marshal(pubData)
	if nil != err {
		global.Log.Error("error ", err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if nil != err {
		return err
	}
	req.SetBasicAuth(global.Conf.Mqtt.Username, global.Conf.Mqtt.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		global.Log.Errorf("Publish.DefaultClient.Do data=%s error=%s", string(data), err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Log.Error("error ReadAll", err)
		return err
	}
	global.Log.Debug("receive resp, ", string(body))
	if resp.StatusCode != 200 {
		global.Log.Error("bad status ", resp.StatusCode)
		return errors.New(resp.Status)
	}
	return nil
}
