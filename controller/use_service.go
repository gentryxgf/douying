package controller

import "douyin/service"

var (
	VideoSer   = service.ServiceGroupApp.VideoService
	MessageSer = service.ServiceGroupApp.MessageService
)
