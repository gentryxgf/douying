package controller

type ControllerGroup struct {
	VideoController         VideoController
	UserRegiterContoller    UserRegiterContoller
	UserVedioListController UserVedioListController
}

var ControllerGroupApp = new(ControllerGroup)
