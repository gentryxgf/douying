package controller

import (
	"douyin/common/global"
	"douyin/models/request"

	"douyin/common/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FollowListController struct{}

func (FollowListController) FollowListView(c *gin.Context) {

	var ur request.FollowListRequest
	if err := c.ShouldBindQuery(&ur); err != nil {
		global.Log.Error("FollowContoller.FollowListView 请求参数错误", zap.Error(err))
	}

	userid, _ := strconv.Atoi(ur.UserId)
	_claim, exist := c.Get("claim")
	claim, _ := _claim.(*jwt.UserClaim)
	if !exist {
		global.Log.Error("UserVedioListController.UserVedioListView  claim 参数错误")
	}
	res, err := FollowListSer.FollowList(int64(userid), claim)

	if err != nil {
		global.Log.Error("FollowController.FollowListView USE FollowSer.FollowList ERROR", zap.Error(err))
		// response.FAIL(&response.FllowResponse{}, response.FOLLOW_LIST_ERROR, c)
		// return
	}
	c.JSON(http.StatusOK, res)
}
