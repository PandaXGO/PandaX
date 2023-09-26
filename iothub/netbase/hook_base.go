package netbase

import (
	"encoding/json"
	"log"
	"pandax/apps/device/services"
	"pandax/iothub/server/emqxserver/protobuf"
	"pandax/pkg/global"
	"pandax/pkg/tool"
	"regexp"
	"strings"
	"time"
)

func Auth(authToken string) bool {
	// 根据token，去查设备Id以及设备类型
	if authToken == "pandax" {
		return true
	}
	etoken := &tool.DeviceAuth{}
	// redis 中有就查询，没有就添加
	exists, err := global.RedisDb.Exists(global.RedisDb.Context(), authToken).Result()
	if exists == 1 {
		err = global.RedisDb.Get(authToken, etoken)
	} else {
		device, err := services.DeviceModelDao.FindOneByToken(authToken)
		log.Println(err)
		if err != nil {
			global.Log.Infof("设备token %s 不存在", authToken)
			return false
		}
		etoken, err = services.GetDeviceToken(device)
		if err != nil {
			global.Log.Infof("设备TOKEN %s添加缓存失败", authToken)
			return false
		}
	}
	if err != nil {
		global.Log.Infof("invalid authToken %s", authToken)
		return false
	}
	// 判断token是否过期了, 设备过期
	if etoken.ExpiredAt < time.Now().Unix() {
		global.Log.Infof("设备authToken %s 失效", authToken)
		return false
	}
	return true
}

// 解析遥测数据类型 返回标准带时间戳格式
func UpdateDeviceTelemetryData(data string) map[string]interface{} {
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

func UpdateDeviceAttributesData(data string) map[string]interface{} {
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

// encode data
func EncodeData(jsonData interface{}) ([]byte, error) {
	byteData, err := json.Marshal(jsonData)
	if nil != err {
		return nil, err
	}
	return byteData, nil
}

func GetRequestIdFromTopic(reg, topic string) (requestId string) {
	re := regexp.MustCompile(reg)
	res := re.FindStringSubmatch(topic)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}
