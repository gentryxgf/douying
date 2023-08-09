package service

type ServiceGroup struct {
	VideoService         VideoService
	UserRegisterService  UserRegisterService
	UserVideoListService UserVideoListService
}

var ServiceGroupApp = new(ServiceGroup)
