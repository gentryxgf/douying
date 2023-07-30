package controller

import "douying/service"

// 返回前端的公共数据
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 登录注册返回前端的数据
type LoginRegisterResponse struct {
	Response
	*service.LoginResponse
}
