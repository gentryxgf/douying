package routers

import (
	"douyin/controller"
	"douyin/middleware"
)

func (router RouterGroup) FollowRouter() {
	app := controller.ControllerGroupApp.FollowContoller
	router.POST("/relation/action", middleware.JwtUser(), app.FollowView)
	router.GET("/relation/follow/list", middleware.JwtUser(), controller.ControllerGroupApp.FollowListContoller.FollowListView)
}
