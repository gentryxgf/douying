package dao

type DaoGroup struct {
	VideoDao VideoDao
}

var DaoGroupApp = new(DaoGroup)
