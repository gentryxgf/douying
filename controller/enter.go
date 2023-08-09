package controller

type ControllerGroup struct {
	VedioController      VedioController
	UserRegiterContoller UserRegiterContoller
}

var ControllerGroupApp = new(ControllerGroup)
