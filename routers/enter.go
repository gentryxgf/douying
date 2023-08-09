package routers

import (
	"douyin/common/global"

	"github.com/gin-gonic/gin"
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
	routerGroupApp.UserRegisterRouter()

	return router
}
