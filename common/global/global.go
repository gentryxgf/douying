package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mini-tiktok/common/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Log    *zap.Logger
)
