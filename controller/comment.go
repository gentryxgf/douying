package controller

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"douyin/models/request"
	"douyin/models/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentController struct{}

func (CommentController) CommentAction(c *gin.Context) {
	var req request.CommentActionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		global.Log.Error("CommentController.CommentAction 请求参数错误", zap.Error(err))
		c.JSON(http.StatusOK, response.CommentActionResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "请求参数错误",
		}})
		return
	}

	// 暂时省去判断视频ID是否存在
	claim, err := jwt.ParseToken(req.Token)
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

	switch req.ActionType {
	case 1:
		if req.CommentText != "" {
			comment, err := CommentSer.AddComment(UserID, req.VideoID, req.CommentText)
			if err != nil {
				c.JSON(http.StatusOK, response.CommentActionResponse{
					Response: response.Response{
						StatusCode: global.SUCCESS,
						StatusMsg:  "评论成功",
					},
					Comment: comment,
				})
			}
		}

	case 2:
		// 省去判断评论ID是否存在
		err := CommentSer.DeleteComment(req.CommentID)
		if err != nil {
			c.JSON(http.StatusOK, response.CommentActionResponse{
				Response: response.Response{
					StatusCode: global.SUCCESS,
					StatusMsg:  "删除评论成功",
				},
			})
		}
	default:
		global.Log.Error("CommentController.CommentAction 请求参数action_type错误")
		c.JSON(http.StatusOK, response.CommentActionResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "请求参数错误",
		}})
		return
	}

	return
}

func (CommentController) CommentList(c *gin.Context) {
	var req request.CommentListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		global.Log.Error("CommentController.CommentList 请求参数错误", zap.Error(err))
		c.JSON(http.StatusOK, response.CommentListResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "请求参数错误",
		}})
		return
	}

	// 暂时省去判断视频ID是否存在

	commentList, err := CommentSer.GetCommentList(req.VideoID)
	if err != nil {
		c.JSON(http.StatusOK, response.CommentListResponse{
			Response: response.Response{
				StatusCode: global.SUCCESS,
				StatusMsg:  "获取评论列表成功",
			},
			CommentList: commentList,
		})
	}

	return
}
