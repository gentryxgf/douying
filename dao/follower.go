package dao

import (
	"douyin/common/global"
	"douyin/models"

	"go.uber.org/zap"
)

type FollowerListDao struct {
}

// 通过ID查询用户的粉丝列表
func (FollowerListDao) FindFollowerListById(id int64) ([]*models.FollowModel, error) {
	var followerList = make([]*models.FollowModel, 0)
	err := global.DB.Model(&models.FollowModel{}).Where("to_user_id = ?", id).Find(&followerList).Error
	if err != nil {
		global.Log.Error("FollowerListDao.FindFollowerListById ERROR ID查询粉丝列表失败", zap.Int64("userid", id), zap.Error(err))
		return nil, err
	}
	return followerList, nil
}

// 通过follow表查询是否是粉丝
func (FollowerListDao) IsFollowByIds(from, to int64) bool {
	var res models.FollowModel
	global.DB.Model(&models.FollowModel{}).Where("from_user_id = ? and to_user_id = ?", from, to).First(&res)
	return res.ID != 0
}
