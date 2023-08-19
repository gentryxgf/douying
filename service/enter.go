package service

type ServiceGroup struct {
	VideoService          VideoService
	UserRegisterService   UserRegisterService
	UserVideoListService  UserVideoListService
	MessageService        MessageService
	FollowService         FollowService
	FollowListService     FollowListService
	FavoriteListService   FavoriteListService
	FavoriteActionService FavoriteActionService //
	CommentService        CommentService
}

var ServiceGroupApp = new(ServiceGroup)
