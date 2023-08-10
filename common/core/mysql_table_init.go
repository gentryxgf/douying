package core

import (
	"douyin/common/global"
	"douyin/models"
	"go.uber.org/zap"
)

func InitMysqlTable() (err error) {
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.VideoModel{},
			&models.CommentModel{},
			&models.FavoriteModel{},
			&models.FollowModel{},
			&models.FriendModel{},
			&models.MessageModel{},
			&models.UserModel{},
			&models.LikeModel{},
		)
	if err != nil {
		global.Log.Error("Mysql生成数据库表结构失败", zap.Error(err))
		return err
	}
	global.Log.Info("Mysql生成数据库表结构成功")
	return nil
}
