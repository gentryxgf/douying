package request

type UploadVideoRequest struct {
	//Data  []byte `json:"data" form:"data" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
}

type VideoFeedRequest struct {
	LatestTime int64 `json:"latest_time" form:"latest_time"`
}
