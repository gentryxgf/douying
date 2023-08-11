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

// 通过follow表查询是否是粉丝
func (UserVideoListDao) IsFollowByIds(from, to int64) bool {
	var res models.FollowModel
	global.DB.Model(&models.FollowModel{}).Where("from_user_id = ? and to_user_id = ?", from, to).First(&res)
	if res.ID == 0 {
		return false
	}
	return true
}

// 通过喜欢表查询是否点赞
func (UserVideoListDao) IsFavoriteByIds(userid, videoid int64) bool {
	var res models.FavoriteModel
	global.DB.Model(&models.FavoriteModel{}).Where("user_id = ? and video_id = ?", userid, videoid).First(&res)
	if res.ID == 0 {
		return false
	}
	return true
}
