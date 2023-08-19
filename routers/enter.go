package routers

import (
	"douyin/common/global"
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

	routerGroup := router.Group("douyin")
	routerGroupApp := RouterGroup{routerGroup}

	routerGroupApp.VideoRouter()
	routerGroupApp.CommentRouter()

	routerGroupApp.UserRegisterRouter()
	routerGroupApp.MessageRouter()

	routerGroupApp.FavoriteListRouter()
	routerGroupApp.FavoriteActionRouter() //

	routerGroupApp.FollowRouter()

	return router
}
