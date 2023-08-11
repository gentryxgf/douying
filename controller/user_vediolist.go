package controller

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"douyin/models/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserVedioListController struct{}

func (UserRegiterContoller) UserVedioListView(c *gin.Context) {
	//获取userid
	var ur request.VideoListRequest
	if err := c.ShouldBindQuery(&ur); err != nil {
		global.Log.Error("UserVedioListController.UserVedioListView 请求参数错误", zap.Error(err))
	}

	userid, _ := strconv.Atoi(ur.UserId)
	_claim, exist := c.Get("claim")
	claim, _ := _claim.(*jwt.UserClaim)
	if !exist {
		global.Log.Error("UserVedioListController.UserVedioListView  claim 参数错误")
	}
	res, err := UserVedioListSer.UserVideoList(int64(userid), claim)
	if err != nil {
		global.Log.Info("UserVedioListController.UserVedioListView USE UserVedioListSer.UserVideoList ERROR 获取视频列表失败", zap.Error(err))
	}
	c.JSON(http.StatusOK, res)
}
