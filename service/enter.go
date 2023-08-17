package service

type ServiceGroup struct {
	VideoService   VideoService
	CommentService CommentService
}

var ServiceGroupApp = new(ServiceGroup)
