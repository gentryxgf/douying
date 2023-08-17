package service

import "douyin/dao"

var (
	VideoDao   = dao.DaoGroupApp.VideoDao
	UserDao    = dao.DaoGroupApp.UserDao
	CommentDao = dao.DaoGroupApp.CommentDao
)
