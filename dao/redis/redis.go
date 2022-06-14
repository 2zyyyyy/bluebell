package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("ping rdb failed.", zap.Error(err))
		return err
	}
	return
}

func Close() {
	err := rdb.Close()
	if err != nil {
		zap.L().Error("redis close failed.", zap.Error(err))
		return
	}
}
