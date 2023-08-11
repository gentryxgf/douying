package routers

import "douyin/controller"

func (router RouterGroup) FavoriteListRouter() {
	app := controller.ControllerGroupApp.FavoriteListController
	router.GET("/favorite/list", app.FavoriteListView)
}
