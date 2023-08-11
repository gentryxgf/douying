package dao

import (
	"douyin/common/global"
	"douyin/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDao struct{}

func (UserDao) GetUserInfo(userID int64) (user models.UserModel, err error) {
	err = global.DB.Where("id = ?", userID).Take(&user).Error
	if err != nil {
		global.Log.Error("UserDao.GetUserInfo USE global.DB.Take, MODEL: UserModel ERROR", zap.Error(err))
	}
	return
}

// IsFollowUser 判断用户A 是否关注 用户B
func (UserDao) IsFollowUser(userAID, userBID int64) (bool, error) {
	var follow models.FollowModel
	err := global.DB.Where("from_user_id = ? and to_user_id = ?", userAID, userBID).Take(&follow).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("UserDao.GetUserInfo USE global.DB.Take, MODEL: FollowModel ERROR", zap.Error(err))
		return false, err
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return true, nil
}
