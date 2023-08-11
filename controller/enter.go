package controller

type ControllerGroup struct {
	VideoController   VideoController
	MessageController MessageController
}

var ControllerGroupApp = new(ControllerGroup)
