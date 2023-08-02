package core

import (
	"context"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"mini-tiktok/common/config"
	"mini-tiktok/common/global"
	"time"
)

func InitRedis(cfg config.RedisConf) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr(),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Log.Error("Redis 连接失败", zap.Error(err))
		return
	}
	global.Redis = rdb
}
