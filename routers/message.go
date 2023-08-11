package routers

import "douyin/controller"

func (router RouterGroup) MessageRouter() {
	app := controller.ControllerGroupApp.MessageController
	router.POST("/message/action", app.SendMessageView)
	router.GET("/message/chat", app.MessageChatView)
}
