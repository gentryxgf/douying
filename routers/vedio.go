package routers

import (
	"douyin/controller"
)

func (router RouterGroup) VideoRouter() {
	app := controller.ControllerGroupApp.VideoController
	router.POST("/publish/action", app.UploadVideoView)
	router.GET("/feed", app.VideoFeedView)
}
