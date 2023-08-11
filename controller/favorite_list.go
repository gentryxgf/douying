package controller

import (
	"douyin/common/global"
	"douyin/models/request"
	"douyin/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type FavoriteListController struct{}

func (FavoriteListController) FavoriteListView(c *gin.Context) {
	// 获取token和title
	var cr request.FavoriteListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error(" FavoriteListRequest参数错误", zap.Error(err))
	}
	/*claim, err := jwt.parseToken(cr.Token)
	if err != nil {
		global.Log.Error("token 错误", zap.Error(err))
		c.JSON(http.StatusOK, response.Response{
		StatusCode: global.ERROR,
			StatusMsg:  "token 错误",
		})
		c.Abort()
		return
	}
	c.Set("claim",claim)


	UserID, _ := c.Get("claim")        //通过token解析得到id
	UserID := cr.FavoriteListRequestId //直接得id?
	*/

	//var UserID int64 = 10 //测试
	UserID := cr.FavoriteListRequestId
	// 获取点赞列表业务处理
	FoundR, err := FavoriteListSer.FavoriteList(UserID) //FoundR 对应文档要返回的结构体
	// 获取点赞列表业务处理
	if err != nil {
		global.Log.Error("controllerfindFavorite ERROR", zap.Error(err))
	}
	if len(FoundR) == 0 {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: http.StatusOK,
			StatusMsg:  "未查询到信息",
		})
		global.Log.Error("FindFavorite is zero", zap.Error(err))
	} else {
		fmt.Println("最终结果")
		fmt.Println(FoundR)
		c.JSON(http.StatusOK, FoundR)
	}

}
