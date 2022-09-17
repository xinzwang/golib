package utils

import (
	"encoding/json"
	"sync"
	"time"

	"e.coding.net/itdesk/weixin/golib/config"

	"github.com/go-redis/redis"
)

var (
	redisConfig config.RedisConfig
	redisLock   sync.Mutex
	redisIns    *redis.Client
)

func SetRedisConfig(cfg *config.RedisConfig) {
	redisConfig = *cfg
}

func NewRedisByConfig() error {
	redisIns = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       0,
	})
	_, err := redisIns.Ping().Result()
	return err
}

func RedisIns() *redis.Client {
	if redisIns == nil {
		redisLock.Lock()
		defer redisLock.Unlock()
		if redisIns == nil {
			NewRedisByConfig()
		}
	}
	return redisIns
}

func GetObject(key string, obj interface{}) error {
	if value, err := RedisIns().Get(key).Bytes(); err != nil {
		return err
	} else {
		_ = json.Unmarshal(value, &obj)
		return nil
	}
}

func SetObject(key string, obj interface{}, expiration time.Duration) error {
	buffer, _ := json.Marshal(obj)
	return RedisIns().Set(key, buffer, expiration).Err()
}
