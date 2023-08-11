package service

type ServiceGroup struct {
	VideoService   VideoService
	MessageService MessageService
}

var ServiceGroupApp = new(ServiceGroup)
