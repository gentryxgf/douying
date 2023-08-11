package dao

type DaoGroup struct {
	UserRegisterDao    UserRegisterDao
	UserVideoListDao   UserVideoListDao
	VideoDao           VideoDao
	UserDao            UserDao
	MessageDao         MessageDao
	FavoriteListDao    FavoriteListDao
	FavoriteActiontDao FavoriteActionDao
}

var DaoGroupApp = new(DaoGroup)
