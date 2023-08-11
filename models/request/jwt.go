package request

type JwtToken struct {
	Token string `json:"token" form:"token" binding:"required"`
}
