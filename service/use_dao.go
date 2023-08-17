package service

import "douyin/dao"

var (
	UserRegisterDao  = dao.DaoGroupApp.UserRegisterDao
	UserVideoListDao = dao.DaoGroupApp.UserVideoListDao
	VideoDao         = dao.DaoGroupApp.VideoDao
	UserDao          = dao.DaoGroupApp.UserDao
	MessageDao       = dao.DaoGroupApp.MessageDao
	FarvoriteListDao = dao.DaoGroupApp.FavoriteListDao //
	CommentDao = dao.DaoGroupApp.CommentDao
)
