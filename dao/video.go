package dao

import (
	"douyin/common/global"
	"douyin/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
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
		global.Log.Error("VideoDao.CreateVideo USE global.DB.Create, MODEL: VideoModel ERROR", zap.Error(err))
		return err
	}

	return nil
}

func (VideoDao) GetVideoList(latestTime string) (list []models.VideoModel, err error) {
	list = make([]models.VideoModel, 0, 30)

	err = global.DB.Where("created_at <= ?", latestTime).Order("created_at desc").Limit(30).Find(&list).Error
	if err != nil {
		global.Log.Error("VideoDao.GetVideoList USE global.DB.Find, MODEL: VideoModel ERROR", zap.Error(err))
		return nil, err
	}

	return
}

// IsLikeVideo 判断用户是否点赞视频
func (VideoDao) IsLikeVideo(userID, videoID int64) (bool, error) {
	var like models.LikeModel
	err := global.DB.Where("user_id = ? and video_id = ?", userID, videoID).Take(&like).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("VideoDao.IsLikeVideo USE global.DB.Take, MODEL: LikeModel ERROR", zap.Error(err))
		return false, err
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return true, nil
}
