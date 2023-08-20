package controller

import "douyin/service"

var (
	VideoSer         = service.ServiceGroupApp.VideoService
	UserRegisterSer  = service.ServiceGroupApp.UserRegisterService
	UserVedioListSer = service.ServiceGroupApp.UserVideoListService
	MessageSer       = service.ServiceGroupApp.MessageService
	FollowSer        = service.ServiceGroupApp.FollowService
	FollowListSer    = service.ServiceGroupApp.FollowListService
	FollowerListSer  = service.ServiceGroupApp.FollowerListService
)
