package request

type FavoriteActionRequest struct {
	Token      string `form:"token"  binding:"required" msg:"Token不能为空"`
	VideoID    int64  `form:"videoid"  binding:"required" msg:"videoid不能为空"`
	ActionType int32  `form:"actiontype"  binding:"required" msg:"actiontype不能为空"`
}
