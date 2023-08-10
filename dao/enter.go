package dao

type DaoGroup struct {
	VideoDao VideoDao
	UserDao  UserDao
}

var DaoGroupApp = new(DaoGroup)
