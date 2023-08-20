package controller

import (
	"douyin/common/global"
	"douyin/common/utils"
	"douyin/models/request"
	"douyin/models/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FollowContoller struct{}

func (FollowContoller) FollowView(c *gin.Context) {
	//获取to_user_id和action_type
	var ur request.FollowRequest
	if err := c.ShouldBindQuery(&ur); err != nil {
		global.Log.Error("FollowContoller.FollowView 请求参数错误", zap.Error(err))
	}

	// 获取用户ID
	claim, err := utils.GetUserID(c)
	if err != nil {
		global.Log.Error("FollowContoller.FollowView 获取UserClaim出错", zap.Error(err))
		response.FAIL(&response.FllowResponse{}, response.SERVER_COMMON_ERROR, c)
		return
	}

	user_id := claim.UserID
	to_user_id := ur.ToUserID
	action_type := ur.ActionType
	err = FollowSer.Follow(user_id, to_user_id, action_type)

	if err != nil {
		global.Log.Error("FollowController.FollowView USE FollowSer.Follow ERROR", zap.Error(err))
		response.FAIL(&response.FllowResponse{}, response.REQUEST_PARAM_ERROR, c)
		return
	}
	response.OK(&response.FllowResponse{}, c)
}
