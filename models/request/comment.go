package request

type CommentActionRequest struct {
	Token       string `json:"token" form:"token"  binding:"required"`
	VideoID     int64  `json:"video_id" form:"video_id" binding:"required"`
	ActionType  int32  `json:"action_type" form:"action_type" binding:"required"`
	CommentText string `json:"comment_text" form:"comment_text"`
	CommentID   int64  `json:"comment_id" form:"comment_id"`
}

type CommentListRequest struct {
	VideoID int64 `json:"video_id" form:"video_id" binding:"required"`
}
