package request

type FollowRequest struct {
	ToUserID   int64 `json:"to_user_id" form:"to_user_id" binding:"required"`
	ActionType int32 `json:"action_type" form:"action_type" binding:"required"`
}
