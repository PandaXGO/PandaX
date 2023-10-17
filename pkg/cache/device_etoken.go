package cache

import (
	"context"
	"github.com/PandaXGO/PandaKit/rediscli"
	"time"
)

var RedisDb *rediscli.RedisDB

// SetDeviceEtoken key 是设备的时候为token， 是子设备的时候为设备编码
func SetDeviceEtoken(key string, value any, duration time.Duration) error {
	return RedisDb.Set(key, value, duration)
}

// GetDeviceEtoken value 是参数指针
func GetDeviceEtoken(key string, value interface{}) error {
	return RedisDb.Get(key, value)
}

// DelDeviceEtoken 删除指定的key
func DelDeviceEtoken(key string) error {
	return RedisDb.Del(context.Background(), key).Err()
}

func ExistsDeviceEtoken(key string) bool {
	exists, _ := RedisDb.Exists(RedisDb.Context(), key).Result()
	if exists == 1 {
		return true
	} else {
		return false
	}
}
