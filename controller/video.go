package controller

import (
	"douyin/common/global"
	"douyin/models/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VedioController struct{}

func (VedioController) UploadVideoView(c *gin.Context) {
	// 获取token和title
	var cr request.UploadVedioRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error("UploadVideoView 参数错误", zap.Error(err))
	}
	video, err := c.FormFile("data")
	if err != nil {
		global.Log.Error("UploadVedioView 获取视频出错", zap.Error(err))
	}

	// 视频上传业务处理
	err = VideoSer.UploadVideo(video)
	if err != nil {
		global.Log.Error("VideoSer.UploadVideo ERROR", zap.Error(err))
	}
}
