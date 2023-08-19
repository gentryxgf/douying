package controller

type ControllerGroup struct {
	VideoController         VideoController
	UserRegisterContoller   UserRegisterContoller
	UserVedioListController UserVedioListController
	MessageController       MessageController
	FollowContoller         FollowContoller
	FollowListContoller     FollowListController
}

var ControllerGroupApp = new(ControllerGroup)
