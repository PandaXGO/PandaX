package iothub

import (
	"encoding/json"
	exhook "pandax/iothub/protobuf"
	"pandax/pkg/global"
	"pandax/pkg/tool"
	"regexp"
	"strings"
	"time"
)

func (s *HookService) auth(username, password string) bool {
	// 根据token，去查设备Id以及设备类型
	if username == "pandax" && password == "pandax" {
		return true
	}
	etoken := &tool.DeviceAuth{}
	err := global.RedisDb.Get(username, etoken)
	if err != nil {
		global.Log.Infof("invalid username %s", username)
		return false
	}
	// 判断token是否过期了, 设备过期
	if etoken.ExpiredAt < time.Now().Unix() {
		global.Log.Infof("设备%s: Token失效", username)
		return false
	}
	if etoken.Token != password {
		global.Log.Infof("invalid password %s", password)
		return false
	}
	return true
}

// 解析遥测数据类型 返回标准带时间戳格式
func updateDeviceTelemetryData(data string) map[string]interface{} {
	tel := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &tel)
	if err != nil {
		global.Log.Error("上传遥测数据结构错误")
		return nil
	}
	if _, ok := tel["ts"]; ok {
		if _, ok := tel["values"]; ok {
			marshal, err := json.Marshal(tel["values"])
			if err != nil {
				global.Log.Error("上传遥测数据values数据结构错误")
				return nil
			}
			resTel := make(map[string]interface{})
			json.Unmarshal(marshal, &resTel)
			resTel["ts"] = tel["ts"]
			return resTel
		}
	}
	tel["ts"] = time.Now().Format("2006-01-02 15:04:05.000")
	return tel
}

func updateDeviceAttributesData(data string) map[string]interface{} {
	tel := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &tel)
	if err != nil {
		global.Log.Error("上传属性数据结构错误")
		return nil
	}
	tel["ts"] = time.Now().Format("2006-01-02 15:04:05.000")
	return tel
}
func GetUserName(Clientinfo *exhook.ClientInfo) string {
	// coap 协议 用户名最大支持5个字符
	protocol := Clientinfo.GetProtocol()
	var username string
	if protocol == "coap" {
		username = Clientinfo.GetClientid()
	} else if protocol == "lwm2m" {
		username = SplitLwm2mClientID(Clientinfo.GetClientid(), 0)
	} else {
		username = Clientinfo.GetUsername()
	}
	return username
}

func SplitLwm2mClientID(lwm2mClientID string, index int) string {
	// LwM2M client id should be username@password
	idArray := strings.Split(lwm2mClientID, "@")
	if len(idArray) < 2 || index > (len(idArray)+1) {
		return ""
	}
	return idArray[index]
}

func GetPassword(Clientinfo *exhook.ClientInfo) string {
	protocol := Clientinfo.GetProtocol()
	var pw string
	if protocol == "lwm2m" {
		pw = SplitLwm2mClientID(Clientinfo.GetClientid(), 1)
	} else {
		pw = Clientinfo.GetPassword()
	}
	return pw
}

// encode data
func EncodeData(jsonData interface{}) ([]byte, error) {
	byteData, err := json.Marshal(jsonData)
	if nil != err {
		return nil, err
	}
	return byteData, nil
}

func getRequestIdFromTopic(reg, topic string) (requestId string) {
	re := regexp.MustCompile(reg)
	res := re.FindStringSubmatch(topic)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}
