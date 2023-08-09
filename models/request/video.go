package request

type UploadVideoRequest struct {
	//Data  []byte `json:"data" form:"data" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
}
