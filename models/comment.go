package models

type CommentModel struct {
	MODEL
	CommentID  int64  `json:"comment_id" gorm:"comment_id;not null"`
	UserID     int64  `json:"user_id" gorm:"column:user_id;not null"`
	VideoID    int64  `json:"video_id" gorm:"column:video_id;not null"`
	Content    string `json:"content" gorm:"column:content"`
	CreateDate string `json:"create_date" gorm:"column:create_date"`
}

func (CommentModel) TableName() string {
	return "comment"
}
