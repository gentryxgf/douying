package response

type UploadVideoResponse struct {
	Response
}

type VideoFeedResponse struct {
	Response
	NextTime  int64
	VideoList []VideoDetail `json:"video_list"`
}

type VideoDetail struct {
	ID             int64  `json:"id"`
	Author         Author `json:"author"`
	PlayUrl        string `json:"play_url"`
	CoverUrl       string `json:"cover_url"`
	FavouriteCount int64  `json:"favourite_count"`
	IsFavourite    bool   `json:"is_favourite"`
	Title          string `json:"title"`
}
