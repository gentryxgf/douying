package routers

import (
	"douyin/controller"
)

func (router RouterGroup) VideoRouter() {
	app := controller.ControllerGroupApp.VedioController
	router.POST("/videos", app.UploadVideoView)
}
