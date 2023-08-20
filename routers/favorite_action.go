package routers

import "douyin/controller"

func (router RouterGroup) FavoriteActionRouter() {
	app := controller.ControllerGroupApp.FavoriteActionController
	router.GET("/favorite/action", app.FavoriteActionView)
}
