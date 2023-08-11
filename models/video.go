package models

type VideoModel struct {
	MODEL
	PlayUrl      string `json:"play_url" gorm:"column:play_url"`
	CoverUrl     string `json:"cover_url" gorm:"column:cover_url"`
	Title        string `json:"title" gorm:"column:title"`
	UserID       int64  `json:"user_id" gorm:"column:user_id;not null"`
	LikeCount    int64  `json:"like_count" gorm:"column:like_count"`
	CollectCount int64  `json:"collect_count" gorm:"column:collect_count"`
	LookCount    int64  `json:"look_count" gorm:"column:look_count"`
	CommentCount int64  `json:"comment_count" gorm:"column:comment_count"`
}

func (VideoModel) TableName() string {
	return "video"
}
