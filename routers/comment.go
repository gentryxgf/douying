package routers

import (
	"douyin/controller"
)

func (router RouterGroup) CommentRouter() {
	app := controller.ControllerGroupApp.CommentController
	router.POST("/comment/action", app.CommentAction)
	router.GET("/comment/list", app.CommentList)
}
