package dao

type DaoGroup struct {
	VideoDao   VideoDao
	UserDao    UserDao
	MessageDao MessageDao
}

var DaoGroupApp = new(DaoGroup)
