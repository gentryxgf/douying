package middleware

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"douyin/models/request"
	"douyin/models/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func JwtUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 绑定token参数
		var cr request.JwtToken
		if err := c.ShouldBind(&cr); err != nil {
			global.Log.Error("未携带token", zap.Error(err))
			c.JSON(http.StatusOK, response.Response{
				StatusCode: global.ERROR,
				StatusMsg:  "未携带token",
			})
			c.Abort()
			return
		}

		// 解析token
		claim, err := jwt.ParseToken(cr.Token)
		if err != nil {
			global.Log.Error("token错误", zap.Error(err))
			c.JSON(http.StatusOK, response.Response{
				StatusCode: global.ERROR,
				StatusMsg:  "token错误",
			})
			c.Abort()
			return
		}

		// 判断是否登出
		// 待实现

		c.Set("claim", claim)
		c.Next()
	}
}
