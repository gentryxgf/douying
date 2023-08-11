package service

import (
	"douyin/common/jwt"
	"douyin/models"
	"douyin/models/response"
	"errors"
)

type UserVideoListService struct{}

func (UserVideoListService) UserVideoList(userid int64, claim *jwt.UserClaim) (*response.UserVideoListResponse, error) {
	videoList, err := userVideoList(userid, claim)
	if err != nil {
		res := &response.UserVideoListResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
		return res, err
	}
	res := &response.UserVideoListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		VideoList: videoList,
	}
	return res, nil
}

func userVideoList(userid int64, claim *jwt.UserClaim) ([]response.VideoData, error) {
	user, err := UserRegisterDao.FindUserById(userid)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}
	if user.ID == 0 {
		return nil, errors.New("该用户不存在")
	}
	var videoList = make([]*models.VideoModel, 0)
	videoList, err = UserVideoListDao.FindVideoListById(userid)
	if err != nil {
		return nil, errors.New("查询用户视频失败")
	}
	if len(videoList) == 0 {
		return nil, errors.New("该用户没有发表视频")
	}

	isFollow := UserVideoListDao.IsFollowByIds(claim.UserID, userid)

	userdata := response.UserData{
		Id:                  (*user).ID,
		Name:                (*user).Username,
		FollowCount:         (*user).FollowCount,
		FollowerCount:       (*user).FollowerCount,
		IsFollow:            isFollow,
		Avatar:              (*user).Avatar,
		BackgroundImage:     (*user).BackgroundImage,
		Signature:           (*user).Signature,
		TotalFavoritedCount: (*user).TotalFavoritedCount,
		WorkCount:           (*user).WorkCount,
		FavoriteCount:       (*user).FavoriteCount,
	}

	var videoResponseList = make([]response.VideoData, 20)
	for _, video := range videoList {
		isFavorite := UserVideoListDao.IsFavoriteByIds((*user).ID, (*video).ID)
		videodata := response.VideoData{
			Id:            (*video).ID,
			Author:        userdata,
			PlayUrl:       (*video).PlayUrl,
			CoverUrl:      (*video).CoverUrl,
			FavoriteCount: (*video).LikeCount,
			CommentCount:  (*video).CommentCount,
			IsFavorite:    isFavorite,
			Title:         (*video).Title,
		}
		videoResponseList = append(videoResponseList, videodata)
	}
	return videoResponseList, nil
}
