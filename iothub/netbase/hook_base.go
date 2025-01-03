package netbase

import (
	"encoding/json"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	exhook "pandax/iothub/server/emqxserver/protobuf"
	"pandax/pkg/cache"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"pandax/pkg/tdengine"
	"pandax/pkg/tool"
	"regexp"
	"strings"
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
		etoken.Protocol = device.Product.ProtocolName
		err = cache.SetDeviceEtoken(authToken, etoken.GetMarshal(), time.Hour*24*365)
		if err != nil {
			global.Log.Infof("认证失败，设备TOKEN %s添加缓存失败", authToken)
			return false
		}
	}
	return true
}

// SubAuth 获取子设备的认证信息
func SubAuth(authToken string, productAuth *model.DeviceAuth) (*model.DeviceAuth, bool) {
	defer func() {
		if Rerr := recover(); Rerr != nil {
			global.Log.Error(Rerr)
			return
		}
	}()
	// 解析认证
	tokens := strings.Split(authToken, "_")
	producId := ""
	name := ""
	if len(tokens) >= 2 {
		producId = tokens[0]
		name = tokens[1]
	} else if len(tokens) == 1 {
		name = tokens[0]
	} else {
		return nil, false
	}

	etoken := &model.DeviceAuth{}
	// redis 中有就查询，没有就添加
	err := cache.GetDeviceEtoken(name, etoken)
	if err == nil {
		return etoken, true
	}
	// 判断子设备 已经创建的子设备
	deviceRes, err := services.DeviceModelDao.FindOneByName(name)
	// 没有设备就要创建子设备
	if err != nil {
		product, err := services.ProductModelDao.FindOne(producId)
		if err != nil {
			return nil, false
		}
		//自动创建设备，
		// 1. 创建设备
		device := entity.Device{
			Name:       name,
			Pid:        producId,
			Alias:      productAuth.Name + "子设备",
			ParentId:   productAuth.DeviceId,
			Gid:        productAuth.DeviceGroup,
			Status:     "0",
			LinkStatus: global.ONLINE,
			LastAt:     time.Now(),
			DeviceType: entity.GATEWAYS_DEVICE,
		}
		device.Id = utils.GenerateID("d")
		device.OrgId = product.OrgId
		device.Owner = product.Owner
		device.Protocol = product.ProtocolName
		deviceD, err := services.DeviceModelDao.Insert(device)
		if err != nil {
			return nil, false
		}
		etoken = services.GetDeviceToken(deviceD)
		etoken.Protocol = product.ProtocolName
	} else {
		etoken = services.GetDeviceToken(&deviceRes.Device)
		etoken.Protocol = deviceRes.Product.ProtocolName
	}
	err = cache.SetDeviceEtoken(name, etoken.GetMarshal(), time.Hour*24*365)
	if err != nil {
		global.Log.Infof("设备 %s添加缓存失败", name)
		return nil, false
	}
	return etoken, true
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
	tel["ts"] = time.Now().Local().Format("2006-01-02 15:04:05.000")
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
	tel["ts"] = time.Now().Local().Format("2006-01-02 15:04:05.000")
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
	ts := time.Now().Local().Format("2006-01-02 15:04:05.000")
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
