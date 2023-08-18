package dao

import (
	"douyin/common/global"
	"douyin/models"

	"go.uber.org/zap"
)

type CommentDao struct{}

func (CommentDao) CreateComment(userID int64, videoID int64, content string) (models.CommentModel, error) {
	comment := models.CommentModel{
		UserID:  userID,
		VideoID: videoID,
		Content: content,
	}
	err := global.DB.Create(&comment).Error
	if err != nil {
		global.Log.Error("CommentDao创建评论失败", zap.Error(err))
	}
	return comment, nil
}

func (CommentDao) DeleteCommentByID(commentID int64) (err error) {
	err = global.DB.Where("id = ?", commentID).Delete(&models.CommentModel{}).Error
	if err != nil {
		global.Log.Error("CommentDao根据评论ID删除评论失败", zap.Error(err))
	}
	return nil
}

func (CommentDao) QueryCommentByVideoID(videoID int64) ([]models.CommentModel, error) {
	var comments []models.CommentModel
	res := global.DB.Where("video_id = ?", videoID).Find(&comments)
	if res.Error != nil {
		global.Log.Error("CommentDao根据视频ID查询评论列表失败", zap.Error(res.Error))
	}
	return comments, nil
}
