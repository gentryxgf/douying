package models

type CommentModel struct {
	MODEL
	UserID  int64  `json:"user_id" gorm:"column:user_id;not null"`
	VideoID int64  `json:"video_id" gorm:"column:video_id;not null"`
	Content string `json:"content" gorm:"column:content"`
}
