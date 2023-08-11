package dao

type DaoGroup struct {
	UserRegisterDao  UserRegisterDao
	UserVideoListDao UserVideoListDao
	VideoDao         VideoDao
	UserDao          UserDao
	MessageDao       MessageDao
}

var DaoGroupApp = new(DaoGroup)
