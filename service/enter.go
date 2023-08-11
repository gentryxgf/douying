package service

type ServiceGroup struct {
	VideoService         VideoService
	UserRegisterService  UserRegisterService
	UserVideoListService UserVideoListService
	MessageService       MessageService
}

var ServiceGroupApp = new(ServiceGroup)
