package controller

import "douyin/service"

var (
	VideoSer   = service.ServiceGroupApp.VideoService
	CommentSer = service.ServiceGroupApp.CommentService
)
