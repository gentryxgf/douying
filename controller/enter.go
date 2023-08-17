package controller

type ControllerGroup struct {
	VideoController          VideoController
	UserRegisterContoller    UserRegisterContoller
	UserVedioListController  UserVedioListController
	MessageController        MessageController
	FavoriteListController   FavoriteListController
	FavoriteActionController FavoriteActionController //
	CommentController CommentController
}

var ControllerGroupApp = new(ControllerGroup)
