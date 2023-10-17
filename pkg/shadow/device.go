package shadow

import (
	"time"
)

// Device 设备结构
type Device struct {
	Name             string                 // 设备名称
	ProductName      string                 // 设备模型名称
	AttributesPoints map[string]DevicePoint // 设备属性点位列表 key 作为属性
	TelemetryPoints  map[string]DevicePoint // 设备遥测点位列表 key 作为属性
	online           bool                   // 在线状态
	updatedAt        time.Time              // 更新时间
}

// DevicePoint 设备点位结构
type DevicePoint struct {
	Name      string      // 点位名称
	Value     interface{} // 点位值
	UpdatedAt time.Time
}

func NewDevice(deviceName string, productName string, attributes, telemetry map[string]DevicePoint) Device {
	return Device{
		Name:             deviceName,
		ProductName:      productName,
		AttributesPoints: attributes,
		TelemetryPoints:  telemetry,
		online:           true,
	}
}

func NewDevicePoint(pointName string, value interface{}) DevicePoint {
	return DevicePoint{
		Name:      pointName,
		Value:     value,
		UpdatedAt: time.Now(),
	}
}
