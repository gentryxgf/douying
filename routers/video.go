package routers

import (
	"douyin/controller"
	"douyin/middleware"
)

func (router RouterGroup) VideoRouter() {
	app := controller.ControllerGroupApp.VideoController
	router.POST("/publish/action", middleware.JwtUser(), app.UploadVideoView)
	router.GET("/feed", app.VideoFeedView)
}
