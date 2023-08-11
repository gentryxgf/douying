package service

import "douyin/dao"

var (
	VideoDao   = dao.DaoGroupApp.VideoDao
	UserDao    = dao.DaoGroupApp.UserDao
	MessageDao = dao.DaoGroupApp.MessageDao
)
