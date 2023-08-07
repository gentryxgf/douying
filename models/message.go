package models

type MessageModel struct {
	MODEL
	SendUserID     int64  `json:"send_user_id" gorm:"column:send_user_id;not null"`
	SendUsername   string `json:"send_username" gorm:"column:send_username"`
	SendUserAvator string `json:"send_user_avator" gorm:"column:send_user_avator"`
	RevUserID      int64  `json:"rev_user_id" gorm:"column:rev_user_id;not null"`
	RevUsername    string `json:"rev_username" gorm:"column:rev_username"`
	RevUserAvator  string `json:"rev_user_avator" gorm:"column:rev_user_avator"`
	IsRead         bool   `json:"is_read" gorm:"column:is_read;default:false"`
	Content        string `json:"content" gorm:"column:content"`
}
