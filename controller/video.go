package controller

import (
	"douyin/common/global"
	"douyin/models/request"
	"douyin/models/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type VideoController struct{}

func (VideoController) UploadVideoView(c *gin.Context) {
	// 获取token和title
	var cr request.UploadVideoRequest
	if err := c.ShouldBind(&cr); err != nil {
		global.Log.Error("VideoController.UploadVideoView 参数错误", zap.Error(err))
		c.JSON(http.StatusOK, response.UploadVideoResponse{Response: response.Response{
			StatusCode: global.ERROR,
			StatusMsg:  "参数错误",
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

	var userID int64 = 10

	// 视频上传业务处理
	err = VideoSer.UploadVideo(video, cr.Title, userID, c)
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
