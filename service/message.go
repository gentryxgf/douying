package service

import (
	"douyin/common/global"
	"douyin/models/response"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MessageService struct{}

func (MessageService) SendMessage(sendUserID, revUserID int64, content string) error {
	// 获取发送用户信息
	sendUser, err := UserDao.GetUserInfo(sendUserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("MessageService.SendMessage USE UserDao.GetUserInfo ERROR", zap.Error(err))
		return err
	} else if err == gorm.ErrRecordNotFound {
		global.Log.Error("查找用户不存在", zap.Int64("userID", revUserID))
		return errors.New("用户不存在")
	}

	// 获取接收用户信息
	revUser, err := UserDao.GetUserInfo(revUserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("MessageService.SendMessage USE UserDao.GetUserInfo ERROR", zap.Error(err))
		return err
	} else if err == gorm.ErrRecordNotFound {
		global.Log.Error("查找用户不存在", zap.Int64("userID", revUserID))
		return errors.New("用户不存在")
	}

	// 保存到数据库
	err = MessageDao.CreateMessage(
		sendUser.ID, sendUser.Username, sendUser.Avatar,
		revUser.ID, revUser.Username, revUser.Avatar,
		content,
	)
	if err != nil {
		global.Log.Error("MessageService.SendMessage USE MessageDao.CreateMessage ERROR", zap.Error(err))
		return err
	}

	return nil
}

func (MessageService) MessageChat(sendUserID, revUserID int64, preMsgTime string) (list []response.Message, err error) {

	// 先判断聊天对象是否存在
	_, err = UserDao.GetUserInfo(revUserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("MessageService.MessageChat USE UserDao.GetUserInfo ERROR", zap.Error(err))
		return
	} else if err == gorm.ErrRecordNotFound {
		global.Log.Error("查找用户不存在", zap.Int64("userID", revUserID))
		return nil, errors.New("用户不存在")
	}

	messageList, err := MessageDao.MessageChat(sendUserID, revUserID, preMsgTime)
	if err != nil {
		global.Log.Error("MessageService.SendMessage USE MessageDao.MessageChat ERROR", zap.Error(err))
		return
	}

	list = make([]response.Message, 0, len(messageList))

	for _, model := range messageList {
		list = append(list, response.Message{
			SendUserID:     model.SendUserID,
			SendUsername:   model.SendUsername,
			SendUserAvatar: model.SendUserAvatar,
			RevUserID:      model.RevUserID,
			RevUsername:    model.RevUsername,
			RevUserAvatar:  model.RevUserAvatar,
			IsRead:         model.IsRead,
			Content:        model.Content,
			CreateTime:     model.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return
}
