package service

type ServiceGroup struct {
	VideoService          VideoService
	UserRegisterService   UserRegisterService
	UserVideoListService  UserVideoListService
	MessageService        MessageService
	FavoriteListService   FavoriteListService
	FavoriteActionService FavoriteActionService //
}

var ServiceGroupApp = new(ServiceGroup)
