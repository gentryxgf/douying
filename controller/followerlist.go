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

type FollowerListController struct{}

func (FollowerListController) FollowerListView(c *gin.Context) {

	var ur request.FollowListRequest
	if err := c.ShouldBindQuery(&ur); err != nil {
		global.Log.Error("FollowerlistContoller.FollowerListView 请求参数错误", zap.Error(err))
	}

	userid, _ := strconv.Atoi(ur.UserId)
	_claim, exist := c.Get("claim")
	claim, _ := _claim.(*jwt.UserClaim)
	if !exist {
		global.Log.Error("FollowerListController.FollowerListView  claim 参数错误")
	}
	res, err := FollowerListSer.FollowerList(int64(userid), claim)

	if err != nil {
		global.Log.Error("FollowerListController.FollowerListView USE FollowerSer.FollowerList ERROR", zap.Error(err))
		// response.FAIL(&response.FllowResponse{}, response.FOLLOW_LIST_ERROR, c)
		// return
	}
	c.JSON(http.StatusOK, res)
}
