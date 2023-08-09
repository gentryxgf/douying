package core

import (
	"douyin/common/config"
	"douyin/common/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitMysql(cfg config.MysqlConf) {
	if cfg.Host == "" {
		global.Log.Warn("未配置MySQL， 取消MySQL连接")
		return
	}
	dsn := cfg.Dsn()
	fmt.Println(dsn)
	// 自定义gorm Logger
	var mysqlLogger logger.Interface
	if cfg.LogLevel == "debug" {
		// 显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 只显示错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		global.Log.Error("MySQL 连接失败", zap.Error(err))
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdelConns) // 连接池最大空闲连接数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) // 连接池最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //数据库连接的最大生存时间
	global.DB = db
}
