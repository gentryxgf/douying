package dao

import (
	"douyin/common/global"
	"douyin/models"
	"go.uber.org/zap"
)

type VideoDao struct{}

func (VideoDao) CreateVideo(userID int64, playUrl, coverUrl, title string) (err error) {

	err = global.DB.Create(&models.VideoModel{
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
		UserID:   userID,
	}).Error
	if err != nil {
		global.Log.Error("VideoDao.CreateVideo USE global.DB.Create ERROR", zap.Error(err))
		return err
	}

	return nil
}
