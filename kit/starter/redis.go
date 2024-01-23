package starter

import (
	"fmt"
	"pandax/kit/logger"

	"github.com/go-redis/redis"
)

func ConnRedis(host, password string, db, port int) *redis.Client {
	// 设置redis客户端
	logger.Log.Infof("连接redis [%s:%d]", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	// 测试连接
	_, e := rdb.Ping().Result()
	if e != nil {
		logger.Log.Panic(fmt.Sprintf("连接redis失败! [%s:%d]", host, port))
	}
	return rdb
}
