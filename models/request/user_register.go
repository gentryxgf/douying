package request

type UserRegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required,max=32"`
	Password string `json:"password" form:"password" binding:"required,min=8,max=32"`
}
