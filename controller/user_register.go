package controller

import (
	"douyin/common/global"
	"douyin/models/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserRegiterContoller struct{}

func (UserRegiterContoller) UserRegisterView(c *gin.Context) {
	//获取username和password
	var ur request.UserRegisterRequest
	if err := c.ShouldBindQuery(&ur); err != nil {
		global.Log.Error("UserRegisterController.UserRegisterView 请求参数错误", zap.Error(err))
	}
	username := ur.Username
	password := ur.Password
	res, err := UserRegisterSer.UserRegister(username, password)
	if err != nil {
		global.Log.Info("UserRegisterController.UserRegisterView USE UserRegisterSer.UserRegister Error 注册失败", zap.Error(err))
	}
	c.JSON(http.StatusOK, res)
}

func (UserRegiterContoller) UserLoginView(c *gin.Context) {
	//获取username和password
	var ur request.UserRegisterRequest
	if err := c.ShouldBindQuery(&ur); err != nil {
		global.Log.Error("UserRegisterController.UserLoginView 请求参数错误", zap.Error(err))
	}
	username := ur.Username
	password := ur.Password
	res, err := UserRegisterSer.UserLogin(username, password)
	if err != nil {
		global.Log.Info("UserRegisterController.UserLoginView USE UserRegisterSer.UserLogin ERROR 登录失败", zap.Error(err))
	}
	c.JSON(http.StatusOK, res)
}
