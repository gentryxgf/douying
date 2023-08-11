package request

type VideoListRequest struct {
	UserId string `json:"user_id" form:"user_id" binding:"required"`
}
