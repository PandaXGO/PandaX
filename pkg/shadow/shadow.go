package shadow

import (
	"errors"
	"fmt"
	"pandax/pkg/global"
	"sync"
	"time"
)

var UnknownDeviceErr = errors.New("unknown device")
var DeviceRepeatErr = errors.New("device already exists")
var BindPointDataErr = errors.New("bind online point data can't be parsed")

type OnlineChangeCallback func(deviceName string, online bool) // 设备上/下线回调

// DeviceShadow 设备影子
type DeviceShadow interface {
	AddDevice(device Device) (err error)
	GetDevice(deviceName string) (device Device, err error)

	SetDevicePoint(deviceName, pointType, pointName string, value interface{}) (err error)
	GetDevicePoint(deviceName, pointType, pointName string) (value DevicePoint, err error)
	GetDevicePoints(deviceName, pointType string) (points map[string]DevicePoint, err error)

	GetDeviceUpdateAt(deviceName string) (time.Time, error)

	GetDeviceStatus(deviceName string) (online bool, err error)

	SetOnline(deviceName string) (err error)
	SetOffline(deviceName string) (err error)

	SetOnlineChangeCallback(handlerFunc OnlineChangeCallback)

	// StopStatusListener 停止设备状态监听
	StopStatusListener()

	// SetDeviceTTL 设备影子过期时间
	SetDeviceTTL(ttl int)
}

type deviceShadow struct {
	m           *sync.Map
	ticker      *time.Ticker
	handlerFunc OnlineChangeCallback
	ttl         int
}

var DeviceShadowInstance DeviceShadow

func init() {
	shadow := &deviceShadow{
		m:      &sync.Map{},
		ticker: time.NewTicker(time.Second),
	}
	go shadow.checkOnOff()
	DeviceShadowInstance = shadow
}

func InitDeviceShadow(deviceName, ProductId string) Device {
	device, err := DeviceShadowInstance.GetDevice(deviceName)
	if err == UnknownDeviceErr {
		attributes := make(map[string]DevicePoint)
		telemetry := make(map[string]DevicePoint)
		device = NewDevice(deviceName, ProductId, attributes, telemetry)
		DeviceShadowInstance.AddDevice(device)
		//shadow.DeviceShadowInstance.SetDeviceTTL()
	}
	return device
}

func (d *deviceShadow) AddDevice(device Device) (err error) {
	if _, ok := d.m.Load(device.Name); ok {
		return DeviceRepeatErr
	}
	device.updatedAt = time.Now()
	d.m.Store(device.Name, device)
	return nil
}

// SetDeviceTTL 设备影子过期时间
func (d *deviceShadow) SetDeviceTTL(ttl int) {
	d.ttl = ttl
}
func (d *deviceShadow) GetDevice(deviceName string) (device Device, err error) {
	if deviceAny, ok := d.m.Load(deviceName); ok {
		return deviceAny.(Device), nil
	} else {
		return Device{}, UnknownDeviceErr
	}
}

func (d *deviceShadow) SetDevicePoint(deviceName, pointType, pointName string, value interface{}) (err error) {
	deviceAny, ok := d.m.Load(deviceName)
	if !ok {
		return UnknownDeviceErr
	}
	device := deviceAny.(Device)
	// update point value
	device.updatedAt = time.Now()

	if pointType == global.TslAttributesType {
		device.AttributesPoints[pointName] = NewDevicePoint(pointName, value)
	} else if pointType == global.TslTelemetryType {
		device.TelemetryPoints[pointName] = NewDevicePoint(pointName, value)
	} else {
		return errors.New("设备属性类型错误")
	}
	// update
	d.m.Store(deviceName, device)
	return
}

func (d *deviceShadow) GetDevicePoint(deviceName, pointType, pointName string) (value DevicePoint, err error) {
	if deviceAny, ok := d.m.Load(deviceName); ok {
		device := deviceAny.(Device)
		if device.online == false || time.Now().Sub(device.updatedAt) > time.Duration(d.ttl)*time.Second {
			return
		}
		if pointType == global.TslAttributesType {
			return device.AttributesPoints[pointName], nil
		} else if pointType == global.TslTelemetryType {
			return device.TelemetryPoints[pointName], nil
		} else {
			return value, errors.New("设备属性类型错误")
		}
	} else {
		return value, UnknownDeviceErr
	}
}

func (d *deviceShadow) GetDevicePoints(deviceName, pointType string) (points map[string]DevicePoint, err error) {
	if deviceAny, ok := d.m.Load(deviceName); ok {
		if pointType == global.TslAttributesType {
			return deviceAny.(Device).AttributesPoints, nil
		} else if pointType == global.TslTelemetryType {
			return deviceAny.(Device).TelemetryPoints, nil
		} else {
			return points, errors.New("设备属性类型错误")
		}
	} else {
		return nil, UnknownDeviceErr
	}
}

func (d *deviceShadow) GetDeviceUpdateAt(deviceName string) (time.Time, error) {
	if deviceAny, ok := d.m.Load(deviceName); ok {
		return deviceAny.(Device).updatedAt, nil
	} else {
		return time.Time{}, UnknownDeviceErr
	}
}

func (d *deviceShadow) changeOnOff(deviceName string, online bool) (err error) {
	if deviceAny, ok := d.m.Load(deviceName); ok {
		device := deviceAny.(Device)
		if device.online != online {
			device.online = online
			device.updatedAt = time.Now()
			d.m.Store(deviceName, device)
			d.handlerCallback(deviceName, online)
		}
	} else {
		return UnknownDeviceErr
	}
	return
}

func (d *deviceShadow) GetDeviceStatus(deviceName string) (online bool, err error) {
	if deviceAny, ok := d.m.Load(deviceName); ok {
		device := deviceAny.(Device)
		return device.online, nil
	} else {
		return false, UnknownDeviceErr
	}
}

func (d *deviceShadow) SetOnline(deviceName string) (err error) {
	return d.changeOnOff(deviceName, true)
}

func (d *deviceShadow) SetOffline(deviceName string) (err error) {
	return d.changeOnOff(deviceName, false)
}

func (d *deviceShadow) SetOnlineChangeCallback(handlerFunc OnlineChangeCallback) {
	d.handlerFunc = handlerFunc
}

func (d *deviceShadow) StopStatusListener() {
	d.ticker.Stop()
}

func (d *deviceShadow) checkOnOff() {
	for range d.ticker.C {
		d.m.Range(func(key, value any) bool {
			if device, ok := value.(Device); ok {
				// fix: when ttl == 0, device always offline
				if d.ttl == 0 {
					return true
				}

				if device.online && time.Now().Sub(device.updatedAt) > time.Duration(d.ttl)*time.Second {
					_ = d.SetOffline(device.Name)
				}
			}
			return true
		})
	}
}

func (d *deviceShadow) handlerCallback(deviceName string, online bool) {
	if d.handlerFunc != nil {
		go d.handlerFunc(deviceName, online)
	}
}

// 解析在离线状态绑定点位值（支持数据类型：int、float、string、bool）
func parseOnlineBindPV(pv interface{}) (online bool, err error) {
	switch v := pv.(type) {
	case string:
		return parseStrOnlineBindPV(v)
	case int8, int16, int32, int, int64, uint8, uint16, uint32, uint, uint64:
		s := fmt.Sprintf("%d", v)
		return parseStrOnlineBindPV(s)
	case float32, float64:
		s := fmt.Sprintf("%.0f", v)
		return parseStrOnlineBindPV(s)
	case bool:
		if v {
			return true, nil
		}
		return
	default:
		return false, BindPointDataErr
	}
}

func parseStrOnlineBindPV(pv string) (online bool, err error) {
	onlineList := []string{"on", "online", "1", "true", "yes"}
	offlineList := []string{"off", "offline", "0", "false", "no"}
	for i, _ := range onlineList {
		if pv == onlineList[i] {
			return true, nil
		}
	}
	for i, _ := range offlineList {
		if pv == offlineList[i] {
			return false, nil
		}
	}
	return false, BindPointDataErr
}
