package model

type UserModel struct {
	MODEL
	Username            string `json:"username" gorm:"column:username"`
	Password            string `json:"password" gorm:"column:password"`
	Avator              string `json:"avator" gorm:"column:avator"`
	BackgroundImage     string `json:"background_image" gorm:"column:background_image"`
	Signature           string `json:"signature" gorm:"column:signature;type:text"`
	FollowCount         int64  `json:"follow_count" gorm:"column:follow_count"`                   // 关注数量
	FollowerCount       int64  `json:"follower_count" gorm:"column:follower_count"`               // 粉丝数量
	TotalFavoritedCount int64  `json:"total_favorited_count" gorm:"column:total_favorited_count"` // 获赞数量
	FavoriteCount       int64  `json:"favorite_count" gorm:"column:favorite_count"`               // 点赞数量
	WorkCount           int64  `json:"work_count" gorm:"column:work_count"`                       // 作品数量
}
