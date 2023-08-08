package service

type ServiceGroup struct {
	VideoService VideoService
}

var ServiceGroupApp = new(ServiceGroup)
