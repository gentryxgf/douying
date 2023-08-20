package service

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"douyin/models"
	"douyin/models/response"
	"errors"

	"go.uber.org/zap"
)

type FollowerListService struct{}

func (FollowerListService) FollowerList(userid int64, claim *jwt.UserClaim) (*response.FollowListResponse, error) {
	followerUserList, err := FollowerUserList(userid, claim)
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
		FollowList: followerUserList,
	}
	return res, nil
}
func FollowerUserList(userid int64, claim *jwt.UserClaim) ([]response.UserData, error) {

	// 先对userid验证是否为注册用户
	user, err := UserRegisterDao.FindUserById(userid)
	if err != nil {
		global.Log.Error("FollowerListService.FollowerList USE UserRegisterDao.FindUserById ERROR", zap.Error(err))
		return nil, errors.New("查询粉丝失败")
	}
	if user.ID == 0 {
		return nil, errors.New("该粉丝不存在")
	}

	// fmt.Println("service:", user)

	var followList = make([]*models.FollowModel, 0, 10)
	// 查询userid的粉丝列表
	followList, err = FollowerListDao.FindFollowerListById(userid)
	if err != nil {
		global.Log.Error("FollowerListService.FollowerUserList USE FollowerListDao.FindFollowerListById ERROR", zap.Error(err))
		return nil, errors.New("查询粉丝列表失败")
	}
	if len(followList) == 0 {
		return nil, errors.New("该用户没有粉丝")
	}
	// fmt.Println("service:", len(followList))

	// // 构造返回前端的数据格式
	var followResponseList = make([]response.UserData, 0, 20)
	for _, followuser := range followList {
		// 先对userid验证是否为注册用户
		user, err := UserRegisterDao.FindUserById(followuser.FromUserID)
		if err != nil {
			global.Log.Error("FollowerListService.FollowerUserList USE UserRegisterDao.FindUserById ERROR", zap.Error(err))
			return nil, errors.New("查询粉丝失败")
		}
		if user.ID == 0 {
			return nil, errors.New("粉丝不存在")
		}
		isFollow := FollowerListDao.IsFollowByIds(claim.UserID, user.ID)
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
