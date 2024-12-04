package shadow

import (
	"errors"
	"fmt"
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
	HasDevice(deviceName string) bool
	DeleteDevice(deviceName ...string) error

	GetDeviceUpdateAt(deviceName string) (time.Time, error)
	GetDeviceStatus(deviceName string) (online bool, err error)
	RefreshDeviceStatus(deviceName string)
	SetOnline(deviceName string) (err error)
	SetOffline(deviceName string) (err error)

	SetOnlineChangeCallback(handlerFunc OnlineChangeCallback)

	// StopStatusListener 停止设备状态监听
	StopStatusListener()

	// SetDeviceTTL 设备影子过期时间, 过期设备将下线
	SetDeviceTTL(ttl int)
}

type deviceShadow struct {
	m           *sync.Map
	ticker      *time.Ticker
	handlerFunc OnlineChangeCallback //上下线执行的回调函数
	ttl         int
}

var DeviceShadowInstance DeviceShadow

func init() {
	shadow := &deviceShadow{
		m:           &sync.Map{},
		ticker:      time.NewTicker(time.Second),
		ttl:         3600, // 默认1小时
		handlerFunc: deviceHandler,
	}
	go shadow.checkOnOff()
	DeviceShadowInstance = shadow
}

func InitDeviceShadow(deviceName string) Device {
	device, err := DeviceShadowInstance.GetDevice(deviceName)
	if err == UnknownDeviceErr {
		device = NewDevice(deviceName)
		DeviceShadowInstance.AddDevice(device)
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

func (d *deviceShadow) HasDevice(deviceName string) bool {
	if _, ok := d.m.Load(deviceName); ok {
		return ok
	}
	return false
}

func (d *deviceShadow) DeleteDevice(deviceName ...string) error {
	if len(deviceName) == 0 {
		return nil
	}
	for _, v := range deviceName {
		d.m.Delete(v)
	}
	return nil
}

func (d *deviceShadow) GetDeviceUpdateAt(deviceName string) (time.Time, error) {
	if deviceAny, ok := d.m.Load(deviceName); ok {
		return deviceAny.(Device).updatedAt, nil
	} else {
		return time.Time{}, UnknownDeviceErr
	}
}

func (d *deviceShadow) changeOnOff(deviceName string, online bool) (err error) {
	deviceAny, ok := d.m.Load(deviceName)
	if !ok {
		return UnknownDeviceErr
	}
	device := deviceAny.(Device)
	if device.online != online {
		device.online = online
		device.updatedAt = time.Now()
		d.m.Store(deviceName, device)
		d.handlerCallback(deviceName, online)
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

// RefreshDeviceStatus 刷新设备状态
func (d *deviceShadow) RefreshDeviceStatus(deviceName string) {
	if deviceAny, ok := d.m.Load(deviceName); ok {
		device := deviceAny.(Device)
		if device.online {
			return
		}
		device.online = true
		device.updatedAt = time.Now()
		d.m.Store(deviceName, device)
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

// 定时检测设备是否离线
func (d *deviceShadow) checkOnOff() {
	for range d.ticker.C {
		d.m.Range(func(key, value interface{}) bool {
			device, ok := value.(Device)
			if !ok {
				return true
			}
			if device.online && time.Since(device.updatedAt) > time.Duration(d.ttl)*time.Second {
				_ = d.SetOffline(device.Name)
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
	for _, v := range onlineList {
		if pv == v {
			return true, nil
		}
	}
	for _, v := range offlineList {
		if pv == v {
			return false, nil
		}
	}
	return false, BindPointDataErr
}
