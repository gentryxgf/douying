package routers

import (
	"douyin/common/global"
	"douyin/controller"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {

	gin.SetMode(global.Config.SystemConf.Env)
	router := gin.Default()
	router.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	router.POST("/videos", controller.ControllerGroupApp.VideoController.UploadVideoView)

	routerGroup := router.Group("douyin")
	routerGroupApp := RouterGroup{routerGroup}

	routerGroupApp.VideoRouter()

	return router
}
