package dao

import (
	"douyin/common/global"
	"douyin/models"

	"go.uber.org/zap"
)

type FollowListDao struct{}

// 通过ID查询用户发布的视频列表
func (FollowListDao) FindFollowListById(id int64) ([]*models.FollowModel, error) {
	var followList = make([]*models.FollowModel, 0)
	err := global.DB.Model(&models.FollowModel{}).Where("from_user_id = ?", id).Find(&followList).Error
	if err != nil {
		global.Log.Error("FollowListDao.FindFollowListById ERROR ID查询用户关注列表失败", zap.Int64("userid", id), zap.Error(err))
		return nil, err
	}
	return followList, nil
}

// 通过follow表查询是否是粉丝
func (FollowListDao) IsFollowByIds(from, to int64) bool {
	var res models.FollowModel
	global.DB.Model(&models.FollowModel{}).Where("from_user_id = ? and to_user_id = ?", from, to).First(&res)
	return res.ID != 0
}
