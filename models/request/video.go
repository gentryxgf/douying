package request

type UploadVedioRequest struct {
	Token string `json:"token" form:"token" binding:"required"`
	//Data  []byte `json:"data" form:"data" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
}
