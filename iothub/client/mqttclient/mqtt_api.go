package mqttclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"pandax/pkg/global"
	"sync"
)

type MqttClient struct {
	sync.Map
}

type MqttConfig struct {
	HttpBroker string
	Username   string
	Password   string
	Qos        int
}

const (
	ClientsInfo         = "client"
	SubscribeTopicsInfo = "subscribe"
)

var Session MqttClient

func GetEmqInfo(infoType string) ([]map[string]interface{}, error) {
	var url string
	switch infoType {
	case ClientsInfo:
		url = fmt.Sprintf("%s/v5/clients?_page=1&_limit=100000", global.Conf.Mqtt.HttpBroker)
	case SubscribeTopicsInfo:
		url = fmt.Sprintf("%s/v5/subscriptions?_page=1&_limit=100000", global.Conf.Mqtt.HttpBroker)
	default:
		return nil, errors.New("invalid infoType")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
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
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		global.Log.Error("body Unmarshal error", err)
		return nil, err
	}
	data, ok := result["data"].([]map[string]interface{})
	if !ok {
		return nil, errors.New("result error")
	}
	return data, nil
}

func Publish(topic, clientId string, payload interface{}) error {
	if clientId == "" {
		return errors.New("未获取到MQTT连接")
	}
	global.Log.Debugf("send data to clientId: %s, topic:%s, payload: %v", clientId, topic, payload)
	url := fmt.Sprintf("%s/v5/publish", global.Conf.Mqtt.HttpBroker)
	pubData := map[string]interface{}{
		"topic":    topic,
		"payload":  payload,
		"qos":      global.Conf.Mqtt.Qos,
		"retain":   false,
		"clientid": clientId,
	}
	data, err := json.Marshal(pubData)
	if err != nil {
		global.Log.Error("error ", err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
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
	if resp.StatusCode != http.StatusOK {
		global.Log.Error("bad status ", resp.StatusCode)
		return errors.New(resp.Status)
	}
	return nil
}
