package dao

import (
	"douyin/common/global"
	"douyin/models"

	"go.uber.org/zap"
)

type UserVideoListDao struct{}

// 通过ID查询用户发布的视频列表
func (UserVideoListDao) FindVideoListById(id int64) ([]*models.VideoModel, error) {
	var videoList = make([]*models.VideoModel, 0)
	err := global.DB.Model(&models.VideoModel{}).Where("user_id = ?", id).Find(&videoList).Error
	if err != nil {
		global.Log.Error("ID查询用户发布视频列表失败", zap.Int64("userid", id), zap.Error(err))
		return nil, err
	}
	return videoList, nil
}
