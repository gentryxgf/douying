package dao

type DaoGroup struct {
	UserRegisterDao  UserRegisterDao
	UserVideoListDao UserVideoListDao
	VideoDao         VideoDao
	UserDao          UserDao
	MessageDao       MessageDao
	FollowDao        FollowDao
	FollowListDao    FollowListDao
	FollowerListDao  FollowerListDao
}

var DaoGroupApp = new(DaoGroup)
