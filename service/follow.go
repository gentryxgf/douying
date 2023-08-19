package service

import (
	"douyin/common/global"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	Follow   int32 = 1
	UnFollow int32 = 2
)

type FollowService struct{}

func (FollowService) Follow(fromUserID, toUserID int64, actionType int32) error {
	// 获取关注用户信息
	_, err := UserDao.GetUserInfo(fromUserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("MessageService.SendMessage USE UserDao.GetUserInfo ERROR", zap.Error(err))
		return err
	} else if err == gorm.ErrRecordNotFound {
		global.Log.Error("查找用户不存在", zap.Int64("userID", fromUserID))
		return errors.New("用户不存在")
	}

	// 获取被关注用户信息
	_, err = UserDao.GetUserInfo(toUserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("MessageService.SendMessage USE UserDao.GetUserInfo ERROR", zap.Error(err))
		return err
	} else if err == gorm.ErrRecordNotFound {
		global.Log.Error("查找用户不存在", zap.Int64("userID", toUserID))
		return errors.New("用户不存在")
	}
	// fmt.Println("service:", fromUser, toUser)
	// 保存到数据库
	if actionType == Follow {
		err = FollowDao.CreateFollow(fromUserID, toUserID)
		if err != nil {
			global.Log.Error("FollowService.Follow USE FollowDao.CreateFollow ERROR", zap.Error(err))
			return err
		}
	} else if actionType == UnFollow {
		err = FollowDao.CancelFollow(fromUserID, toUserID)
		if err != nil {
			global.Log.Error("FollowService.Follow USE FollowDao.CancelFollow ERROR", zap.Error(err))
			return err
		}
	}

	return err
}
