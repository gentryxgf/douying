package routers

import (
	"douyin/controller"
	"douyin/middleware"
)

func (router RouterGroup) UserRegisterRouter() {
	app := controller.ControllerGroupApp.UserRegiterContoller
	router.POST("/user/register", app.UserRegisterView)
	router.POST("/user/login", app.UserLoginView)
	router.GET("/publish/list", middleware.JwtUser(), app.UserVedioListView)
}
