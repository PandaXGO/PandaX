package shadow

import (
	"pandax/apps/device/services"
	"pandax/pkg/global"
	"time"
)

// Device 设备结构
type Device struct {
	Name      string    // 设备名称
	online    bool      // 在线状态
	updatedAt time.Time // 更新时间
}

func NewDevice(deviceName string) Device {
	return Device{
		Name:      deviceName,
		online:    false,
		updatedAt: time.Now(),
	}
}

func deviceHandler(deviceName string, online bool) {
	if online {
		services.DeviceModelDao.UpdateStatus(deviceName, global.ONLINE)
	} else {
		services.DeviceModelDao.UpdateStatus(deviceName, global.OFFLINE)
	}
}
