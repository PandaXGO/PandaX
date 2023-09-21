package shadow

import "time"

// Device 设备结构
type Device struct {
	Name            string         // 设备名称
	ProductName     string         // 设备模型名称
	Points          map[string]any // 设备点位列表 key参数名 value 值
	online          bool           // 在线状态
	disconnectTimes int            // 断开连接次数，60秒内超过3次判定离线
	updatedAt       time.Time      // 更新时间
}

func NewDevice(deviceName string, productName string, points map[string]any) Device {
	return Device{
		Name:        deviceName,
		ProductName: productName,
		Points:      points,
		online:      true,
	}
}
