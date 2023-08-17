package service

import (
	"douyin/common/global"
	"douyin/dao"
	"fmt"
	"go.uber.org/zap"
)

type FavoriteActionService struct{}

func (FavoriteActionService) FavoriteActionClick(UserID int64, VideoID int64, ActionType int32) error {

	if ActionType == 1 { //点赞操作
		b, err := dao.DaoGroupApp.FavoriteActiontDao.QueryFavoriteBYyUVid(UserID, VideoID)
		if err != nil {
			global.Log.Error("service 点赞前先查询失败", zap.Error(err))
			return err
		}
		if b {
			fmt.Println("已经点过赞")

		} else {
			err := dao.DaoGroupApp.FavoriteActiontDao.CreateFavorite(UserID, VideoID)
			if err != nil {
				global.Log.Error("service 点赞失败", zap.Error(err))
				return err
			}
		}
	} else if ActionType == 2 { //取消点赞
		b, err := dao.DaoGroupApp.FavoriteActiontDao.QueryFavoriteBYyUVid(UserID, VideoID)
		if err != nil {
			global.Log.Error("service 点赞前先查询失败", zap.Error(err))
			return err
		}
		if !b {
			fmt.Println("没点过赞")

		} else {
			err := dao.DaoGroupApp.FavoriteActiontDao.DeleteFavorite(UserID, VideoID)
			if err != nil {
				global.Log.Error("service 取消点赞失败", zap.Error(err))
				return err
			}
		}
	}
	return nil

}
