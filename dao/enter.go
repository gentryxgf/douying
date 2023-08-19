package dao

type DaoGroup struct {
	UserRegisterDao    UserRegisterDao
	UserVideoListDao   UserVideoListDao
	VideoDao           VideoDao
	UserDao            UserDao
	MessageDao         MessageDao
	FavoriteListDao    FavoriteListDao
	FavoriteActiontDao FavoriteActionDao //
	CommentDao         CommentDao
	FollowDao          FollowDao
	FollowListDao      FollowListDao
}

var DaoGroupApp = new(DaoGroup)
