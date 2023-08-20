package service

import (
	"douyin/common/global"
	"douyin/dao"
	"douyin/models/response"

	"go.uber.org/zap"
)

type CommentService struct{}

func (CommentService) AddComment(userID int64, videoID int64, content string) (response.Comment, error) {
	comment, err := dao.DaoGroupApp.CommentDao.CreateComment(userID, videoID, content)
	if err != nil {
		global.Log.Error("添加评论失败", zap.Error(err))
	}
	author, err := dao.DaoGroupApp.FavoriteListDao.QueryUserByVideoID(videoID)
	if err != nil {
		global.Log.Error("根据视频ID查找作者失败", zap.Error(err))
	}
	user := GetCommentUser(userID, author.ID)
	return response.Comment{
		ID:         comment.ID,
		User:       user,
		Content:    comment.Content,
		CreateDate: comment.CreatedAt.Format("01-02"),
	}, nil
}

func (CommentService) DeleteComment(commentID int64) error {
	return dao.DaoGroupApp.CommentDao.DeleteCommentByID(commentID)
}

func (CommentService) GetCommentList(videoID int64) ([]response.Comment, error) {
	var commentList []response.Comment
	comments, err := dao.DaoGroupApp.CommentDao.QueryCommentByVideoID(videoID)
	if err != nil {
		global.Log.Error("获取评论列表失败", zap.Error(err))
	}
	for _, comment := range comments {
		author, err := dao.DaoGroupApp.FavoriteListDao.QueryUserByVideoID(videoID)
		if err != nil {
			global.Log.Error("根据视频ID查找作者失败", zap.Error(err))
		}
		user := GetCommentUser(comment.UserID, author.ID)
		commentList = append(commentList, response.Comment{
			ID:         comment.ID,
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}
	return commentList, nil
}

func GetCommentUser(userID int64, authorID int64) response.Author {
	userInfo, err := dao.DaoGroupApp.UserDao.GetUserInfo(userID)
	if err != nil {
		global.Log.Error("UserDao.GetUserInfo USE global.DB.Take, MODEL: UserModel ERROR", zap.Error(err))
	}
	user := response.Author{
		ID:              userInfo.ID,
		Name:            userInfo.Username,
		FollowCount:     userInfo.FollowCount,
		FollowerCount:   userInfo.FollowerCount,
		Avatar:          userInfo.Avatar,
		BackgroundImage: userInfo.BackgroundImage,
		Signature:       userInfo.Signature,
		TotalFavorited:  userInfo.TotalFavoritedCount,
		WorkCount:       userInfo.WorkCount,
		FavouriteCount:  userInfo.FavoriteCount,
	}
	isFollow, err := UserDao.IsFollowUser(userID, authorID)
	if err != nil {
		user.IsFollow = false
	} else {
		user.IsFollow = isFollow
	}
	return user
}
