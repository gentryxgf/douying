package dao

import (
	"douyin/common/global"
	"douyin/models"
	"errors"

	"go.uber.org/zap"
)

type FollowDao struct{}

func (q *FollowDao) CreateFollow(fromUid int64, toUid int64) (err error) {

	err = q.SelectFollow(fromUid, toUid)
	if err == nil {
		global.Log.Warn("FollowDao.CreateFollow :Duplicate follow", zap.Int64("fromUid", fromUid), zap.Int64("toUid", toUid))
		return errors.New("重复关注")
	}
	err = global.DB.Create(&models.FollowModel{
		FromUserID: fromUid,
		ToUserID:   toUid,
	}).Error
	if err != nil {
		global.Log.Error("FollowDao.CreateFollow USE global.DB.Create, MODEL: FollowModel ERROR", zap.Error(err))
	}
	return err
}

func (q *FollowDao) CancelFollow(fromUid int64, toUid int64) (err error) {
	err = q.SelectFollow(fromUid, toUid)
	if err != nil {
		global.Log.Error("FollowDao.CancelFollow :No follow record", zap.Int64("fromUid", fromUid), zap.Int64("toUid", toUid))
		return errors.New("未找到关注记录")
	}
	err = global.DB.Where("from_user_id = ? and to_user_id = ?", fromUid, toUid).Delete(&models.FollowModel{}).Error
	if err != nil {
		global.Log.Error("FollowDao.CancelFollow USE global.DB.Delete, MODEL: FollowModel ERROR", zap.Error(err))
	}
	return err
}
func (FollowDao) SelectFollow(fromUid int64, toUid int64) (err error) {

	var follow models.FollowModel
	errQuery := global.DB.Where("from_user_id = ? and to_user_id = ?", fromUid, toUid).First(&follow).Error

	// if errQuery != nil {
	// 	// 检查 ErrRecordNotFound 错误
	// 	if errors.Is(errQuery, gorm.ErrRecordNotFound) {
	// 		// global.Log.Error("UserRegisterDao.FindUserByNameAndPass ERROR 登录查询用户失败", zap.String("name", name), zap.Error(err))
	// 		fmt.Println("Dao Select no dulpicate")
	// 	} else {
	// 		// global.Log.Error("UserRegisterDao.FindUserByNameAndPass ERROR 登录查询用户失败", zap.String("name", name), zap.Error(err))
	// 		// return  err
	// 		fmt.Println("Dao Select error")
	// 	}
	// } else {
	// 	fmt.Println("Dao Select successfully")
	// }
	return errQuery
}
