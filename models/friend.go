package models

type FriendModel struct {
	MODEL
	UserAID int64 `json:"user_a_id" gorm:"column:user_a_id"`
	UserBID int64 `json:"user_b_id" gorm:"column:user_b_id"`
}

func (FriendModel) TableName() string {
	return "friend"
}
