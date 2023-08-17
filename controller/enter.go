package controller

type ControllerGroup struct {
	VideoController VideoController
	CommentController CommentController
}

var ControllerGroupApp = new(ControllerGroup)
