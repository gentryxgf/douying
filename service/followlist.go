package service

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"douyin/models"
	"douyin/models/response"
	"errors"

	"go.uber.org/zap"
)

type FollowListService struct{}

func (FollowListService) FollowList(userid int64, claim *jwt.UserClaim) (*response.FollowListResponse, error) {
	followUserList, err := FollowUserList(userid, claim)
	if err != nil {
		res := &response.FollowListResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
		return res, err
	}
	res := &response.FollowListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		FollowList: followUserList,
	}
	return res, nil
}
func FollowUserList(userid int64, claim *jwt.UserClaim) ([]response.UserData, error) {

	// 先对userid验证是否为注册用户
	user, err := UserRegisterDao.FindUserById(userid)
	if err != nil {
		global.Log.Error("FollowListService.FollowList USE UserRegisterDao.FindUserById ERROR", zap.Error(err))
		return nil, errors.New("查询用户失败")
	}
	if user.ID == 0 {
		return nil, errors.New("该用户不存在")
	}

	// fmt.Println("service:", user)

	var followList = make([]*models.FollowModel, 0, 10)
	// 查询userid发布过的视频列表
	followList, err = FollowListDao.FindFollowListById(userid)
	if err != nil {
		global.Log.Error("FollowListService.FollowUserList USE FollowListDao.FindFollowListById ERROR", zap.Error(err))
		return nil, errors.New("查询关注列表失败")
	}
	if len(followList) == 0 {
		return nil, errors.New("该用户未关注任何人")
	}
	// fmt.Println("service:", len(followList))

	// // 构造返回前端的数据格式
	var followResponseList = make([]response.UserData, 0, 20)
	for _, followuser := range followList {
		// 先对userid验证是否为注册用户
		user, err := UserRegisterDao.FindUserById(followuser.ToUserID)
		if err != nil {
			global.Log.Error("FollowListService.FollowUserList USE UserRegisterDao.FindUserById ERROR", zap.Error(err))
			return nil, errors.New("查询关注的用户失败")
		}
		if user.ID == 0 {
			return nil, errors.New("关注用户不存在")
		}
		isFollow := FollowListDao.IsFollowByIds(claim.UserID, user.ID)
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

		followResponseList = append(followResponseList, userdata)
	}
	return followResponseList, nil
}
