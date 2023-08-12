package dao

import (
	"douyin/common/global"
	"douyin/models"
	"fmt"
)

type FavoriteListDao struct{}

func (FavoriteListDao) QueryVideoIdByUserId(UserId int64) ([]int64, error) { //查询喜欢表 保留视频id的切片

	//定义接收查询结果的结构体变量
	var DaoFavoriteID []int64
	var DaoFavorite []models.FavoriteModel                           //保存查询到的结果结构体
	err := global.DB.Where("user_id = ?", UserId).Find(&DaoFavorite) //按照时间排序 不会写
	if err.Error != nil {
		fmt.Println("dao查询喜欢列表返回视频id切片失败", err)
	}
	for i := 0; i < len(DaoFavorite); i++ { //从结构体里面视频id
		DaoFavoriteID = append(DaoFavoriteID, DaoFavorite[i].VideoID)
	}
	return DaoFavoriteID, nil
}

// 此处实现l一个根据视频id 返回视频结构体

func (FavoriteListDao) QueryVideoByVideoID(VideoID int64) (models.VideoModel, error) {
	var DaoVideoS models.VideoModel
	err := global.DB.Where("id = ?", VideoID).Take(&DaoVideoS)
	if err.Error != nil {
		fmt.Println("dao查询喜欢列表返回视频id失败", err)
	}
	return DaoVideoS, nil
}

// 此处实现了一个根据视频id 返回用户结构体

func (FavoriteListDao) QueryUserByVideoID(VideoID int64) (models.UserModel, error) {
	var DaoUserS models.UserModel
	var DaoVideoS models.VideoModel

	err := global.DB.Where("id = ?", VideoID).Take(&DaoVideoS)
	if err.Error != nil {
		fmt.Println("dao查询喜欢列表根据vid返回vstr失败", err)
	}
	fmt.Println(DaoVideoS)
	userid := DaoVideoS.UserID
	err1 := global.DB.Where("id = ?", userid).Take(&DaoUserS)
	if err.Error != nil {
		fmt.Println("dao查询喜欢列表根据uid返回ustr失败", err1)
	}
	return DaoUserS, nil
}
