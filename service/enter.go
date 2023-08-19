package service

type ServiceGroup struct {
	VideoService         VideoService
	UserRegisterService  UserRegisterService
	UserVideoListService UserVideoListService
	MessageService       MessageService
	FollowService        FollowService
	FollowListService    FollowListService
}

var ServiceGroupApp = new(ServiceGroup)
