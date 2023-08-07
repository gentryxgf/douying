package routers

import (
	"github.com/gin-gonic/gin"
	"mini-tiktok/common/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.SystemConf.Env)
	router := gin.Default()

	routerGroup := router.Group("douyin")
	routerGroupApp := RouterGroup{routerGroup}

	routerGroupApp.VedioRouter()

	return router
}
