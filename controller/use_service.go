package controller

import "douyin/service"

var (
	VideoSer          = service.ServiceGroupApp.VideoService
	UserRegisterSer   = service.ServiceGroupApp.UserRegisterService
	UserVedioListSer  = service.ServiceGroupApp.UserVideoListService
	MessageSer        = service.ServiceGroupApp.MessageService
	FavoriteListSer   = service.ServiceGroupApp.FavoriteListService
	FavoriteActionSer = service.ServiceGroupApp.FavoriteActionService
)
