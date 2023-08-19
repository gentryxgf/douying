package controller

type ControllerGroup struct {
	VideoController          VideoController
	UserRegisterContoller    UserRegisterContoller
	UserVedioListController  UserVedioListController
	MessageController        MessageController
	FollowContoller          FollowContoller
	FollowListContoller      FollowListController
	FavoriteListController   FavoriteListController
	FavoriteActionController FavoriteActionController
	CommentController        CommentController
}

var ControllerGroupApp = new(ControllerGroup)
