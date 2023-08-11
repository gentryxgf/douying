package models

type FavoriteModel struct {
	MODEL
	UserID  int64 `json:"user_id" gorm:"column:user_id;not null"`
	VideoID int64 `json:"video_id" gorm:"column:video_id;not null"`
}

func (FavoriteModel) TableName() string {
	return "favorite"
}
