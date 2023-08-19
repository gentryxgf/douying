package response

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

var message map[int32]string

// 成功返回
const SUCCESS int32 = 0

// 全局错误码
const (
	SERVER_COMMON_ERROR int32 = 100001 + iota
	REQUEST_PARAM_ERROR
	TOKEN_INVALID_ERROR
	TOKEN_GENERATE_ERROR
	DB_MYSQL_ERROR
	DB_REDIS_ERROR
)

// 视频服务错误码
const (
	VIDEO_FILE_PARAM_ERROR int32 = 200001 + iota
	VIDEO_UPLOAD_ERROR
	VIDEO_FEED_ERROR
)

// 消息服务错误码
const (
	MESSAGE_SEND_ERROR int32 = 300001 + iota
	MESSAGE_CHAT_ERROR
)

// 关注服务错误码
const (
	FOLLOW_LIST_ERROR int32 = 400001 + iota
)

func init() {
	message = make(map[int32]string)

	// 全局错误信息
	message[SUCCESS] = "成功"
	message[SERVER_COMMON_ERROR] = "服务器繁忙"
	message[REQUEST_PARAM_ERROR] = "请求参数错误"
	message[TOKEN_INVALID_ERROR] = "无效的token"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"

	// 视频服务错误码
	message[VIDEO_FILE_PARAM_ERROR] = "获取上传视频文件失败"
	message[VIDEO_UPLOAD_ERROR] = "上传视频失败"
	message[VIDEO_FEED_ERROR] = "获取视频流失败"

	// 消息服务错误码
	message[MESSAGE_SEND_ERROR] = "发送消息失败"
	message[MESSAGE_CHAT_ERROR] = "获取消息记录失败"

	// 关注服务错误码
	message[FOLLOW_LIST_ERROR] = "获取关注列表失败"
}

func FAIL(data interface{}, code int32, c *gin.Context) {
	var msg string
	var ok bool
	if msg, ok = message[code]; !ok {
		code = SERVER_COMMON_ERROR
		msg = message[SERVER_COMMON_ERROR]
	}

	value := reflect.ValueOf(data).Elem()
	statusCode := value.FieldByName("StatusCode")
	statusCode.SetInt(int64(code))
	statusMsg := value.FieldByName("StatusMsg")
	statusMsg.SetString(msg)

	c.JSON(http.StatusOK, data)
}

func OK(data interface{}, c *gin.Context) {
	value := reflect.ValueOf(data).Elem()
	statusCode := value.FieldByName("StatusCode")
	statusCode.SetInt(int64(SUCCESS))
	statusMsg := value.FieldByName("StatusMsg")
	statusMsg.SetString(message[SUCCESS])

	c.JSON(http.StatusOK, data)
}
