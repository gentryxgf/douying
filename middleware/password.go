package middleware

import (
	"douyin/common/encrypt"

	"github.com/gin-gonic/gin"
)

const (
	MaxPasswordLength = 32
	MinPasswordLength = 8
)

// 使用MD5算法对传入的密码进行校验以及加密
func MD5Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.Query("password")
		// 对密码进行校验
		// 如果为空或者不符合指定长度，则返回空字符串
		// 如果符合，则对其进行MD5加密返回
		if password == "" || len(password) < MinPasswordLength || len(password) > MaxPasswordLength {
			c.PostForm("password")
		} else {
			c.Set("password", encrypt.Md5([]byte(password)))
		}

		c.Next()
	}
}
