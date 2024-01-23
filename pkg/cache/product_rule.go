package cache

import (
	"pandax/kit/cache"
	"time"
)

var ProductCache = cache.NewTimedCache(cache.NoExpiration, 24*time.Hour)

func ComputeIfAbsentProductRule(key string, fun func(any) (any, error)) (any, error) {
	return ProductCache.ComputeIfAbsent(key, fun)
}

func DelProductRule(key string) {
	ProductCache.Delete(key)
}

func PutProductRule(key string, data any) {
	ProductCache.Put(key, data)
}
