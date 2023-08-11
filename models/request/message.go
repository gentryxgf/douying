package request

type SendMessageRequest struct {
	ToUserID   int64  `json:"to_user_id" form:"to_user_id" binding:"required"`
	ActionType int32  `json:"action_type" form:"action_type" binding:"required"`
	Content    string `json:"content" form:"content" binding:"required"`
}

type MessageChatRequest struct {
	ToUserID   int64 `json:"to_user_id" form:"to_user_id" binding:"required"`
	PreMsgTime int64 `json:"pre_msg_time" form:"pre_msg_time" binding:"required"`
}
