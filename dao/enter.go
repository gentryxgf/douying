package dao

type DaoGroup struct {
	UserRegisterDao  UserRegisterDao
	UserVideoListDao UserVideoListDao
	VideoDao         VideoDao
	UserDao          UserDao
}

var DaoGroupApp = new(DaoGroup)
