package controller

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"douyin/models/request"
	"fmt"
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
		global.Log.Error("UserRegisterView 参数错误", zap.Error(err))
	}

	userid, _ := strconv.Atoi(ur.UserId)
	fmt.Println("userid", userid)
	claim, exist := c.Get("claim")
	x, _ := claim.(*jwt.UserClaim)
	if !exist {
		global.Log.Error("UserVedioListView  claim 参数错误")
	}
	res, err := UserVedioListSer.UserVideoList(int64(userid), x)
	if err != nil {
		global.Log.Info("获取视频列表失败", zap.Error(err))
	}
	c.JSON(http.StatusOK, res)
}
