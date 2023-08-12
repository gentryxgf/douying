package controller

type ControllerGroup struct {
	VideoController         VideoController
	UserRegisterContoller   UserRegisterContoller
	UserVedioListController UserVedioListController
	MessageController       MessageController
}

var ControllerGroupApp = new(ControllerGroup)
