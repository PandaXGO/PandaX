package netbase

import (
	"encoding/json"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	exhook "pandax/iothub/server/emqxserver/protobuf"
	"pandax/kit/utils"
	"pandax/pkg/cache"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"pandax/pkg/tdengine"
	"pandax/pkg/tool"
	"regexp"
	"strings"
	"sync"
	"time"
)

func Auth(authToken string) bool {
	// 根据token，去查设备Id以及设备类型
	if authToken == global.Conf.Mqtt.Username {
		return true
	}
	etoken := &model.DeviceAuth{}
	// redis 中有就查询，没有就添加
	exists := cache.ExistsDeviceEtoken(authToken)
	if exists {
		err := cache.GetDeviceEtoken(authToken, etoken)
		if err != nil {
			global.Log.Infof("认证失败，缓存读取设备错误： %s", err)
			return false
		}
	} else {
		device, err := services.DeviceModelDao.FindOneByToken(authToken)
		if err != nil {
			global.Log.Infof("认证失败，设备token %s 不存在", authToken)
			return false
		}
		etoken = services.GetDeviceToken(&device.Device)
		etoken.DeviceProtocol = device.Product.ProtocolName
		err = cache.SetDeviceEtoken(authToken, etoken.GetMarshal(), time.Hour*24*365)
		if err != nil {
			global.Log.Infof("认证失败，设备TOKEN %s添加缓存失败", authToken)
			return false
		}
	}
	// 判断token是否过期了, 设备过期
	if etoken.ExpiredAt < time.Now().Unix() {
		global.Log.Infof("设备authToken %s 失效", authToken)
		return false
	}
	return true
}

// SubAuth 获取子设备的认证信息
func SubAuth(name string) (*model.DeviceAuth, bool) {
	etoken := &model.DeviceAuth{}
	// redis 中有就查询，没有就添加
	exists := cache.ExistsDeviceEtoken(name)
	if exists {
		err := etoken.GetDeviceToken(name)
		if err != nil {
			global.Log.Infof("认证失败，缓存读取设备错误,无效的设备标识： %s", err)
			return nil, false
		}
	} else {
		device, err := services.DeviceModelDao.FindOneByName(name)
		// 没有设备就要创建子设备
		if err != nil {
			global.Log.Infof("设备标识 %s 不存在, ", name)
			return nil, false
		}
		etoken = services.GetDeviceToken(&device.Device)
		etoken.DeviceProtocol = device.Product.ProtocolName
		err = cache.SetDeviceEtoken(name, etoken.GetMarshal(), time.Hour*24*365)
		if err != nil {
			global.Log.Infof("设备标识 %s添加缓存失败", name)
			return nil, false
		}
	}
	return etoken, true
}

// CreateSubTableField 添加子设备字段
func CreateSubTableField(productId, ty string, fields map[string]interface{}) {
	var group sync.WaitGroup
	for key, value := range fields {
		group.Add(1)
		go func(key string, value any) {
			if key == "ts" {
				return
			}
			check := cache.CheckSubDeviceField(key)
			if !check {
				interfaceType := tool.GetInterfaceType(value)
				err := global.TdDb.AddSTableField(productId+"_"+ty, key, interfaceType, 0)
				if err != nil {
					return
				}
				tsl := entity.ProductTemplate{}
				tsl.Pid = productId
				tsl.Id = utils.GenerateID()
				tsl.Name = key
				tsl.Type = interfaceType
				tsl.Key = key
				tsl.Classify = ty
				// 向产品tsl中添加模型
				services.ProductTemplateModelDao.Insert(tsl)
				cache.SetSubDeviceField(key)
			}
		}(key, value)
	}
	group.Wait()
}

// UpdateDeviceTelemetryData 解析遥测数据类型 返回标准带时间戳格式
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

			format := tool.TimeToFormat(tel["ts"])
			if format == "" {
				return nil
			}
			resTel["ts"] = format
			return resTel
		}
	}
	tel["ts"] = time.Now().Format("2006-01-02 15:04:05.000")
	return tel
}

// UpdateDeviceAttributesData 解析属性数据类型 返回标准带时间戳格式
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

// GetUserName 根据emqx连接信息获取用户名
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

func GetRequestIdFromTopic(reg, topic string) (requestId string) {
	re := regexp.MustCompile(reg)
	res := re.FindStringSubmatch(topic)
	if len(res) > 2 {
		return res[2]
	}
	return ""
}

func GetEventFromTopic(reg, topic string) (identifier string) {
	re := regexp.MustCompile(reg)
	res := re.FindStringSubmatch(topic)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}

// eventType 事件类型 info alarm
func CreateEventInfo(msgType, eventType, content string, deviceAuth *model.DeviceAuth) *DeviceEventInfo {
	ts := time.Now().Format("2006-01-02 15:04:05.000")
	ci := &tdengine.Events{
		DeviceId: deviceAuth.DeviceId,
		Name:     msgType,
		Type:     eventType,
		Content:  content,
		Ts:       ts,
	}
	v, err := json.Marshal(*ci)
	if nil != err {
		return nil
	}
	// 添加设备上线记录
	return &DeviceEventInfo{
		DeviceId:   deviceAuth.DeviceId,
		DeviceAuth: deviceAuth,
		Datas:      string(v),
		Type:       msgType,
	}
}
