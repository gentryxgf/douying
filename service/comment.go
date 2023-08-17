package service

import (
	"douyin/models/response"
	"errors"
)

type CommentService struct{}

func (CommentService) AddComment(videoID int64, content string) (response.Comment, error) {
	var a response.Comment
	return a, nil
}

func (CommentService) DeleteComment(videoID int64, commentID int64) error {
	return errors.New("...")
}

func (CommentService) GetCommentList(videoID int64) ([]response.Comment, error) {
	var commentList []response.Comment
	comments := CommentDao.GetComments(videoID)
	
	for comment := range comments {
		commentList = append(commentList, response.Comment{
			ID: comment.CommentID         
			User: 
			Content    
			CreateDate 
		})
	}
	return nil, nil
}
