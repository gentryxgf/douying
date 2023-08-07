package models

type CommentModel struct {
	MODEL
	UserID  int64  `json:"user_id" gorm:"column:user_id;not null"`
	VedioID int64  `json:"vedio_id" gorm:"column:vedio_id;not null"`
	Content string `json:"content" gorm:"column:content"`
}
