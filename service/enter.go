package service

type ServiceGroup struct {
	VideoService         VideoService
	UserRegisterService  UserRegisterService
	UserVideoListService UserVideoListService
	MessageService       MessageService
	FollowService        FollowService
	FollowListService    FollowListService
	FollowerListService  FollowerListService
}

var ServiceGroupApp = new(ServiceGroup)
