package service

import "douyin/dao"

var (
	UserRegisterDao  = dao.DaoGroupApp.UserRegisterDao
	UserVideoListDao = dao.DaoGroupApp.UserVideoListDao
	VideoDao         = dao.DaoGroupApp.VideoDao
	UserDao          = dao.DaoGroupApp.UserDao
)
