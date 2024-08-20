package cache

import (
	"pandax/kit/cache"
	"strings"
	"time"
)

var SubDeviceField = cache.NewTimedCache(cache.NoExpiration, 24*time.Hour)
var SUBDEVICEKEY = "SUBDEVICEKEY"

func CheckSubDeviceField(field string) bool {
	fields, bool := SubDeviceField.Get(SUBDEVICEKEY)
	if !bool {
		return false
	}
	if !strings.Contains(fields.(string), field) {
		return false
	}
	return true
}

func SetSubDeviceField(data string) {
	fields, bool := SubDeviceField.Get(SUBDEVICEKEY)
	if !bool {
		fields = data
	} else {
		fields = fields.(string) + "," + data
	}
	ProductCache.Put(SUBDEVICEKEY, fields)
}
