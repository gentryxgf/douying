package service

import (
	"douyin/dao"
	"fmt"

	"douyin/models/response"
	"net/http"
)

type FavoriteListService struct{}

// 业务处理 组装数据
func (FavoriteListService) FavoriteList(UserID int64) ([]response.FavoriteListResponse, error) {
	var Favorite []response.FavoriteListResponse //定义变量接受组合后的数据

	VideoId, err := dao.DaoGroupApp.FavoriteListDao.QueryVideoIdByUserId(UserID) //返回视频id切片
	if err != nil {
		fmt.Println("service查询喜欢列表根据vid返回vstr失败", err)
	}
	for _, value := range VideoId {
		VideoStruct, err := dao.DaoGroupApp.FavoriteListDao.QueryVideoByVideoID(value) //返回视频结构体
		if err != nil {
			fmt.Println("service查询喜欢列表根据vid返回vstr失败", err)
		}
		UserStruct, err := dao.DaoGroupApp.FavoriteListDao.QueryUserByVideoID(value) //返回所属用户结构体
		if err != nil {
			fmt.Println("service查询喜欢列表根据vid返回ustr失败", err)
		}
		//向最终结果结构体中组装数据
		var ResultStr response.FavoriteListResponse //定义单个结果结构体
		ResultStr.Response.StatusCode = http.StatusOK
		ResultStr.Response.StatusMsg = "查询喜欢列表成功"
		ResultStr.VideoModel = VideoStruct
		ResultStr.UserModel = UserStruct

		Favorite = append(Favorite, ResultStr) //合成切片
		//fmt.Println(len(Favorite))
		//fmt.Println((Favorite))
		//不加显示未使用
	}
	fmt.Println("最终结果")
	fmt.Println(len(Favorite))
	fmt.Println(Favorite)
	return Favorite, nil
}
