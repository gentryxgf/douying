package main

import (
	"douyin/common/core"
	"douyin/common/global"
	"douyin/routers"
	"go.uber.org/zap"
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

	// 初始化路由
	r := routers.InitRouter()
	addr := global.Config.SystemConf.Addr()
	err := r.Run(addr)
	if err != nil {
		global.Log.Fatal("路由失败", zap.Error(err))
	}
}
