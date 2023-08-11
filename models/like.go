package models

// LikeModel 用户点赞视频表
type LikeModel struct {
	MODEL
	UserID  int64 `json:"user_id" gorm:"column:user_id;not null"`
	VideoID int64 `json:"video_id" gorm:"column:video_id;not null"`
}

func (LikeModel) TableName() string {
	return "like"
}
