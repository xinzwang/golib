package redis

import (
	"sync"

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
