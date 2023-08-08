package controller

type ControllerGroup struct {
	VideoController VideoController
}

var ControllerGroupApp = new(ControllerGroup)
