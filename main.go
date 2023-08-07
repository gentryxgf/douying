package main

import (
	"mini-tiktok/common/core"
	"mini-tiktok/common/global"
)

func main() {
	// 初始化配置文件
	core.InitConf()
	// 初始化日志库
	core.InitZap(global.Config.ZapConf)
	// 初始化mysql
	core.InitMysql(global.Config.MysqlConf)
	// 初始化redis
	core.InitRedis(global.Config.RedisConf)
}
