package controller

import (
	"douying/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	raw, _ := c.Get("password")
	password, _ := raw.(string)

	userLoginResponse, err := service.UserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, LoginRegisterResponse{
			Response: Response{StatusCode: 2, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, LoginRegisterResponse{
		Response:      Response{StatusCode: 0, StatusMsg: "登录成功"},
		LoginResponse: userLoginResponse,
	})
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	raw, _ := c.Get("password")
	password, _ := raw.(string)

	userRegisterResponse, err := service.UserRegister(username, password)
	if err != nil {
		c.JSON(http.StatusOK, LoginRegisterResponse{
			Response: Response{StatusCode: 2, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, LoginRegisterResponse{
		Response:      Response{StatusCode: 0, StatusMsg: "注册成功"},
		LoginResponse: userRegisterResponse,
	})
}

/* func Login(c *gin.Context) {
	username := c.Query("username")
	row, _ := c.Get("password")
	password, ok := row.(string)
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Password parse failed!"},
		})
	}
	token := username + password
	//验证登录信息是否在缓存内
	if mapuser, ok := model.UserLoginInfo[token]; ok {
		//如果在缓存内，验证登录密码是否于缓存中的密码相同
		//若不相同，则返回错误信息
		//若相同，则签发token
		if model.UserLoginInfo[token].Password != password {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Password is wrong!"},
			})
		} else {
			tokenStr, _ := middleware.GenerateToken(username, password)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0, StatusMsg: "Login success!"},
				UserId:   mapuser.ID,
				Token:    tokenStr,
			})
		}
	} else {
		//若不在缓存内，先去数据库查询是否有这个用户
		//如果没有则返回
		//如果有则添加进缓存，并返回token
		var user *model.User
		var err error
		user, err = model.NewUserDaoInstance().QueryUserByNameAndPassword(username, password)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 2, StatusMsg: "Username or Password is wrong!"},
			})
		}
		model.UserLoginInfo[token] = *user
		tokenStr, _ := middleware.GenerateToken(username, password)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Login success!"},
			UserId:   user.ID,
			Token:    tokenStr,
		})
	}
} */
