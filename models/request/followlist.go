package request

type FollowListRequest struct {
	UserId string `json:"user_id" form:"user_id" binding:"required"`
}
