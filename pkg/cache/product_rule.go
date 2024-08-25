package cache

import (
	"github.com/PandaXGO/PandaKit/cache"
	"time"
)

var ProductCache = cache.NewTimedCache(cache.NoExpiration, 24*time.Hour)

func ComputeIfAbsentProductRule(key string, fun func(any) (any, error)) (any, error) {
	return ProductCache.ComputeIfAbsent(key, fun)
}

func GetProductRule(key string) (any, bool) {
	return ProductCache.Get(key)
}

func DelProductRule(key string) {
	ProductCache.Delete(key)
}

func PutProductRule(key string, data any) {
	ProductCache.Put(key, data)
}
