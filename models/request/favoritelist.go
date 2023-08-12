package request

type FavoriteListRequest struct {
	Token string `form:"token"  binding:"required" msg:"Token不能为空"`

	FavoriteListRequestId int64 `form:"id"  binding:"required" msg:"id不能为空"`
}
