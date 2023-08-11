package models

type FollowModel struct {
	MODEL
	FromUserID   int64 `json:"from_user_id" gorm:"column:from_user_id;not null"`
	ToUserID     int64 `json:"to_user_id" gorm:"column:to_user_id;not null"`
	IsFollowBack int   `json:"is_follow_back" gorm:"column:is_follow_back;type:tinyint(1);comment:是否互相关注"`
}

func (FollowModel) TableName() string {
	return "follow"
}
