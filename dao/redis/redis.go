package redis

import (
	"fmt"
	"webapp-scaffold/settings"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			settings.Config.RedisConfig.Host,
			settings.Config.RedisConfig.Port),
		Password: settings.Config.RedisConfig.Password,
		DB:       settings.Config.RedisConfig.DB,
		PoolSize: settings.Config.RedisConfig.PoolSize,
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
