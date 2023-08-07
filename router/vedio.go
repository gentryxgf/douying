package router

import "mini-tiktok/controller"

func (router RouterGroup) VedioRouter() {
	app := controller.ControllerGroupApp.VedioController
	router.POST("/vedio/upload", app.UploadVedioView)
}
