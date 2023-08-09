package routers

import "douyin/controller"

func (router RouterGroup) UserRegisterRouter() {
	app := controller.ControllerGroupApp.UserRegiterContoller
	router.POST("/user/register", app.UserRegisterView)
	router.POST("/user/login", app.UserLoginView)
}
