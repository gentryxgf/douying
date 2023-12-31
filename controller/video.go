package controller

import (
	"douyin/common/global"
	"douyin/common/utils"
	"douyin/models/request"
	"douyin/models/response"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VideoController struct{}

func (VideoController) UploadVideoView(c *gin.Context) {
	// 获取token和title
	var cr request.UploadVideoRequest
	if err := c.ShouldBind(&cr); err != nil {
		global.Log.Error("VideoController.UploadVideoView 请求参数错误", zap.Error(err))
		response.FAIL(&response.UploadVideoResponse{}, response.REQUEST_PARAM_ERROR, c)
		return
	}

	video, err := c.FormFile("data")
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView 获取视频出错", zap.Error(err))
		response.FAIL(&response.UploadVideoResponse{}, response.VIDEO_FILE_PARAM_ERROR, c)
		return
	}

	// 获取用户ID
	claim, err := utils.GetUserID(c)
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView 获取UserClaim出错", zap.Error(err))
		response.FAIL(&response.UploadVideoResponse{}, response.SERVER_COMMON_ERROR, c)
		return
	}

	// 视频上传业务处理
	err = VideoSer.UploadVideo(video, cr.Title, claim.UserID, c)
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView USE VideoSer.UploadVideo ERROR", zap.Error(err))
		response.FAIL(&response.UploadVideoResponse{}, response.VIDEO_UPLOAD_ERROR, c)
		return
	}

	response.OK(&response.UploadVideoResponse{}, c)
}

func (VideoController) VideoFeedView(c *gin.Context) {
	var cr request.VideoFeedRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error("VideoController.VideoFeedView 请求参数错误", zap.Error(err))
		response.FAIL(&response.VideoFeedResponse{}, response.REQUEST_PARAM_ERROR, c)
		return
	}

	latestTime := time.Now().Format("2006-01-02 15:04:05")
	// 判断latestTime是否为空
	if cr.LatestTime != 0 {
		latestTime = time.Unix(cr.LatestTime, 0).Format("2006-01-02 15:04:05")
	}

	/*claim, err := utils.GetUserID(c)
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView 获取UserClaim出错", zap.Error(err))
		response.FAIL(&response.VideoFeedResponse{}, response.SERVER_COMMON_ERROR, c)
		return
	}*/

	nextTime, videoList, err := VideoSer.VideoFeed(latestTime, 4)
	if err != nil {
		global.Log.Error("VideoController.VideoFeedView USE VideoSer.VideoFeed ERROR", zap.Error(err))
		response.FAIL(&response.VideoFeedResponse{}, response.VIDEO_FEED_ERROR, c)
		return
	}

	response.OK(&response.VideoFeedResponse{
		Response: response.Response{
			StatusCode: global.SUCCESS,
			StatusMsg:  "获取视频流成功",
		},
		NextTime:  nextTime,
		VideoList: videoList,
	}, c)
}
