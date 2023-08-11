package service

import (
	"douyin/common/jwt"
	"douyin/models"
	"douyin/models/response"
	"errors"
)

type UserVideoListService struct{}

// 用户的视频发布列表，直接列出用户所有投稿过的视频
// 对返回的数据进行打包
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

// 获取用户的视频投稿列表
func userVideoList(userid int64, claim *jwt.UserClaim) ([]response.VideoData, error) {
	// 先对userid验证是否为注册用户
	user, err := UserRegisterDao.FindUserById(userid)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}
	if user.ID == 0 {
		return nil, errors.New("该用户不存在")
	}
	var videoList = make([]*models.VideoModel, 0, 10)
	// 查询userid发布过的视频列表
	videoList, err = UserVideoListDao.FindVideoListById(userid)
	if err != nil {
		return nil, errors.New("查询用户视频失败")
	}
	if len(videoList) == 0 {
		return nil, errors.New("该用户没有发表视频")
	}
	// token用户是否是userid用户的粉丝
	isFollow := UserVideoListDao.IsFollowByIds(claim.UserID, userid)
	// userid也就是视频作者的信息
	// 因为查询的是指定userid的所有视频列表，只需要将用户信息嵌入进视频列表信息
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
	// 构造返回前端的数据格式
	var videoResponseList = make([]response.VideoData, 0, 20)
	for _, video := range videoList {
		// 判断token用户是否对该视频进行过点赞
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
