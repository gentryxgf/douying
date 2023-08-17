package dao

import (
	"douyin/common/global"
	"douyin/models"

	"go.uber.org/zap"
)

type FavoriteActionDao struct{}

func (FavoriteActionDao) CreateFavorite(userID int64, videoID int64) (err error) { //创建一个favorite 记录
	err = global.DB.Create(&models.FavoriteModel{
		UserID:  userID,
		VideoID: videoID,
	}).Error

	if err != nil {
		global.Log.Error("FavoriteActionDao.CreateFavorite USE global.DB.Create ERROR", zap.Error(err))
		return err
	}
	return nil
}
func (FavoriteActionDao) DeleteFavorite(UserID int64, VideoID int64) (err error) { //删除一个favorite 记录

	err = global.DB.Where("user_id = ?&&video_id = ?", UserID, VideoID).Delete(&models.FavoriteModel{}).Error //
	if err != nil {
		global.Log.Error("FavoriteActionDao.DeleteFavorite USE global.DB.Delete ERROR", zap.Error(err))
	}
	return nil
}
func (FavoriteActionDao) QueryFavoriteBYyUVid(UserID int64, VideoID int64) (b bool, err error) { //查询表中是否有值
	var DaoFavorite []models.FavoriteModel
	err = global.DB.Where("user_id = ?&&video_id = ?", UserID, VideoID).Take(&DaoFavorite).Error //
	if err != nil {
		global.Log.Error("dao.favoriteaction查询favorite失败", zap.Error(err))
		return true, err
	}
	if len(DaoFavorite) != 0 {
		return true, nil
	} else {
		return false, nil
	}

}
