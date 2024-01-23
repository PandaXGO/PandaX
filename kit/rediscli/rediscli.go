package rediscli

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisDB struct {
	*redis.Client
}

func NewRedisClient(host, password string, port, db int) (*RedisDB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &RedisDB{
		Client: rdb,
	}, nil
}

func (rdb *RedisDB) HGet(key, field string, obj interface{}) error {
	return rdb.Client.HGet(context.Background(), key, field).Scan(obj)
}

func (rdb *RedisDB) HSet(key, field string, val interface{}) error {
	return rdb.Client.HSet(context.Background(), key, field, val).Err()
}

func (rdb *RedisDB) HDel(key string, fields ...string) error {
	return rdb.Client.HDel(context.Background(), key, fields...).Err()
}

func (rdb *RedisDB) Get(key string, obj interface{}) error {
	return rdb.Client.Get(context.Background(), key).Scan(obj)
}

func (rdb *RedisDB) GetResult(key string) (string, error) {
	return rdb.Client.Get(context.Background(), key).Result()
}

func (rdb *RedisDB) Set(key string, val interface{}, expir time.Duration) error {
	return rdb.Client.Set(context.Background(), key, val, expir).Err()
}
