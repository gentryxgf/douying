package controller

type ControllerGroup struct {
	VideoController         VideoController
	UserRegisterContoller   UserRegisterContoller
	UserVedioListController UserVedioListController
	MessageController       MessageController
	FollowContoller         FollowContoller
	FollowListContoller     FollowListController
	FollowerListContoller   FollowerListController
}

var ControllerGroupApp = new(ControllerGroup)
