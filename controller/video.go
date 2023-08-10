package controller

import (
	"douyin/common/global"
	"douyin/models/request"
	"douyin/models/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type VideoController struct{}

func (VideoController) UploadVideoView(c *gin.Context) {
	// 获取token和title
	var cr request.UploadVideoRequest
	if err := c.ShouldBind(&cr); err != nil {
		global.Log.Error("VideoController.UploadVideoView 请求参数错误", zap.Error(err))
		c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "请求参数错误",
		}})
		return
	}

	video, err := c.FormFile("data")
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView 获取视频出错", zap.Error(err))
		c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "获取视频文件失败",
		}})
		return
	}

	// 获取用户ID
	/*claim, err := utils.GetUserID(c)
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView 获取UserClaim出错", zap.Error(err))
		c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "系统错误",
		}})
		return
	}*/

	// 视频上传业务处理
	err = VideoSer.UploadVideo(video, cr.Title, 10, c)
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView USE VideoSer.UploadVideo ERROR", zap.Error(err))
		c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "上传视频失败",
		}})
		return
	}

	c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
		StatusCode: global.SUCCESS,
		StatusMsg:  "上传视频成功",
	}})
}

func (VideoController) VideoFeedView(c *gin.Context) {
	var cr request.VideoFeedRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error("VideoController.VideoFeedView 请求参数错误", zap.Error(err))
		c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "请求参数错误",
		}})
		return
	}

	// 判断latestTime是否为空
	if cr.LatestTime == "" {
		cr.LatestTime = time.Now().Format("2006-01-02 15:04:05")
	}

	// 获取用户ID
	/*claim, err := utils.GetUserID(c)
	if err != nil {
		global.Log.Error("VideoController.VideoFeedView 获取UserClaim出错", zap.Error(err))
		c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "系统错误",
		}})
		return
	}*/

	nextTime, videoList, err := VideoSer.VideoFeed(cr.LatestTime, 4)
	if err != nil {
		global.Log.Error("VideoController.VideoFeedView USE VideoSer.VideoFeed ERROR", zap.Error(err))
		c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "获取视频流出错",
		}})
		return
	}

	c.JSON(http.StatusOK, response.VideoFeedResponse{
		Response: response.Response{
			StatusCode: global.SUCCESS,
			StatusMsg:  "获取视频流成功",
		},
		NextTime:  nextTime,
		VideoList: videoList,
	})
}
