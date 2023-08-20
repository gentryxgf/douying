package response

import "douyin/models"

type FavoriteListResponse struct {
	Response
	models.VideoModel
	models.UserModel
}
