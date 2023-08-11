package controller

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"douyin/models/request"
	"douyin/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type FavoriteActionController struct{}

func (FavoriteActionController) FavoriteActionView(c *gin.Context) {
	// 获取token和title
	var cr request.FavoriteActionRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error(" FavoriteListRequest参数错误", zap.Error(err))
	}
	/*{
		var structt = jwt.PayLoad{}
		structt.Username = "ycy"
		structt.UserID = 1
		str, err := jwt.GenToken(structt)
		if err != nil {
			global.Log.Error("生成token出错", zap.Error(err))
		}
		fmt.Println("token")
		fmt.Println(str)
		result, err := jwt.ParseToken(str)
		if err != nil {
			global.Log.Error("解析token出错", zap.Error(err))
		}
		fmt.Println(result.PayLoad)
	}
	*///生成tkoen用来测试

	claim, err := jwt.ParseToken(cr.Token)
	if err != nil {
		global.Log.Error("token 错误", zap.Error(err))
		c.JSON(http.StatusOK, response.Response{
			StatusCode: http.StatusOK, ////////改了
			StatusMsg:  "token 错误",
		})
		c.Abort()
		return
	}
	c.Set("claim", claim)
	ClaimResult, ok := c.Get("claim") //通过token解析得到id
	if !ok {
		fmt.Println("跨中间件取值失败")
	}

	ClaimResultTpye := ClaimResult.(*jwt.UserClaim)
	UserID := ClaimResultTpye.PayLoad.UserID
	VideoID := cr.VideoID
	AtcionType := cr.ActionType
	
	//点赞业务处理
	errr := FavoriteActionSer.FavoriteActionClick(UserID, VideoID, AtcionType) //
	if errr != nil {
		global.Log.Error("controller FavoriteActionclick ERROR", zap.Error(errr))
	} else {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: http.StatusOK,
			StatusMsg:  "操作成功",
		})
	}

}
