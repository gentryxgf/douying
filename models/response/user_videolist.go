package response

type UserVideoListResponse struct {
	Response
	VideoList []VideoData `json:"video_list"`
}

type VideoData struct {
	Id            int64    `json:"id"`
	Author        UserData `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

type UserData struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	Avatar              string `json:"avatar"`
	BackgroundImage     string `json:"background_image" `
	Signature           string `json:"signature" `
	FollowCount         int64  `json:"follow_count"`          // 关注数量
	FollowerCount       int64  `json:"follower_count" `       // 粉丝数量
	TotalFavoritedCount int64  `json:"total_favorited_count"` // 获赞数量
	FavoriteCount       int64  `json:"favorite_count"`        // 点赞数量
	WorkCount           int64  `json:"work_count"`
	IsFollow            bool   `json:"is_follow"`
}
