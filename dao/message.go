package dao

import (
	"douyin/common/global"
	"douyin/models"
	"go.uber.org/zap"
)

type MessageDao struct{}

func (MessageDao) CreateMessage(sendUid int64, sendUName, sendUAvatar string, revUid int64, revUName, revUAvatar, content string) (err error) {
	err = global.DB.Create(&models.MessageModel{
		SendUserID:     sendUid,
		SendUsername:   sendUName,
		SendUserAvatar: sendUAvatar,
		RevUserID:      revUid,
		RevUsername:    revUName,
		RevUserAvatar:  revUAvatar,
		Content:        content,
	}).Error
	if err != nil {
		global.Log.Error("MessageDao.CreateMessage USE global.DB.Create, MODEL: MessageModel ERROR", zap.Error(err))
	}

	return
}

func (MessageDao) MessageChat(sendUid, revUid int64, preMsgTime string) (list []models.MessageModel, err error) {
	list = make([]models.MessageModel, 0)
	err = global.DB.Where("send_user_id + rev_user_id = ? and created_at <= ?", sendUid+revUid, preMsgTime).
		Order("created_at desc").Find(&list).Error
	if err != nil {
		global.Log.Error("MessageDao.MessageChat USE global.DB.Find, MODEL: MessageModel ERROR", zap.Error(err))
	}
	return
}
