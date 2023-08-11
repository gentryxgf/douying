package utils

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (*jwt.UserClaim, error) {
	_claim, ok := c.Get("claim")
	if !ok {
		global.Log.Error("token获取claim失败")
		return nil, errors.New("token获取claim失败")
	}
	claim := _claim.(*jwt.UserClaim)
	return claim, nil
}
