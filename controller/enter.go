package controller

type ControllerGroup struct {
	VideoController         VideoController
	UserRegiterContoller    UserRegiterContoller
	UserVedioListController UserVedioListController
	MessageController       MessageController
}

var ControllerGroupApp = new(ControllerGroup)
