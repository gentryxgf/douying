package routers

import (
	"douyin/controller"
)

func (router RouterGroup) VideoRouter() {
	app := controller.ControllerGroupApp.VideoController
	router.POST("/videos", app.UploadVideoView)
}
