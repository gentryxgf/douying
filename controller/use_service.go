package controller

import "douyin/service"

var (
	VideoSer        = service.ServiceGroupApp.VideoService
	UserRegisterSer = service.ServiceGroupApp.UserRegisterService
)
