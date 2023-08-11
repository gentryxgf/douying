package controller

import (
	"douyin/common/global"
	"douyin/models/request"
	"douyin/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type MessageController struct{}

func (MessageController) SendMessageView(c *gin.Context) {
	var cr request.SendMessageRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error("MessageController.SendMessageView 请求参数错误", zap.Error(err))
		response.FAIL(&response.SendMessageResponse{}, response.REQUEST_PARAM_ERROR, c)
		return
	}

	/*claim, err := utils.GetUserID(c)
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView 获取UserClaim出错", zap.Error(err))
		response.FAIL(&response.SendMessageResponse{}, response.SERVER_COMMON_ERROR, c)
		return
	}*/

	// 业务处理
	err := MessageSer.SendMessage(4, cr.ToUserID, cr.Content)
	if err != nil {
		global.Log.Error("MessageController.SendMessageView USE MessageSer.SendMessage ERROR", zap.Error(err))
		response.FAIL(&response.SendMessageResponse{}, response.REQUEST_PARAM_ERROR, c)
		return
	}

	response.OK(&response.SendMessageResponse{}, c)
}

func (MessageController) MessageChatView(c *gin.Context) {
	var cr request.MessageChatRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error("MessageController.MessageChatView 请求参数错误", zap.Error(err))
		response.FAIL(&response.MessageChatResponse{}, response.REQUEST_PARAM_ERROR, c)
		return
	}

	/*claim, err := utils.GetUserID(c)
	if err != nil {
		global.Log.Error("VideoController.UploadVideoView 获取UserClaim出错", zap.Error(err))
		response.FAIL(&response.MessageChatResponse{}, response.SERVER_COMMON_ERROR, c)
		return
	}*/

	preMsgTime := time.Now().Format("2006-01-02 15:04:05")
	// 判断latestTime是否为空
	if cr.PreMsgTime > 1 {
		preMsgTime = time.Unix(cr.PreMsgTime, 0).Format("2006-01-02 15:04:05")
	}

	fmt.Println(preMsgTime)

	list, err := MessageSer.MessageChat(4, cr.ToUserID, preMsgTime)
	if err != nil {
		global.Log.Error("MessageController.MessageChatView USE MessageSer.MessageChat ERROR", zap.Error(err))
		response.FAIL(&response.MessageChatResponse{}, response.MESSAGE_CHAT_ERROR, c)
		return
	}

	response.OK(&response.MessageChatResponse{MessageList: list}, c)
}
