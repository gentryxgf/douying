package dao

type DaoGroup struct {
	UserRegisterDao  UserRegisterDao
	UserVideoListDao UserVideoListDao
	VideoDao         VideoDao
}

var DaoGroupApp = new(DaoGroup)
