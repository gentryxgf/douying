package service

import (
	"douyin/common/global"
	"douyin/common/jwt"
	"douyin/models/response"
	"errors"

	"go.uber.org/zap"
)

const (
	MaxUsernameLength = 32
	MinPasswordLength = 8
	MaxPasswordLength = 32
)

type UserRegisterService struct{}

type LoginRegisterFlow struct {
	username string
	password string
	data     *response.UserRegisterResponse
	userid   int64
	token    string
}

// 注册服务
func (UserRegisterService) UserRegister(username, password string) (*response.UserRegisterResponse, error) {
	return NewLoginRegisterFlow(username, password).Do("register")
}

// 登录服务
func (UserRegisterService) UserLogin(username, password string) (*response.UserRegisterResponse, error) {
	return NewLoginRegisterFlow(username, password).Do("login")
}

func NewLoginRegisterFlow(username, password string) *LoginRegisterFlow {
	return &LoginRegisterFlow{username: username, password: password}
}

func (q *LoginRegisterFlow) Do(method string) (*response.UserRegisterResponse, error) {
	if err := q.checkParams(); err != nil {
		q.data = &response.UserRegisterResponse{
			Response: response.Response{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			}}
		return q.data, err
	}
	if method == "register" {
		if err := q.registerFlow(); err != nil {
			q.data = &response.UserRegisterResponse{
				Response: response.Response{
					StatusCode: 3,
					StatusMsg:  err.Error(),
				}}
			return q.data, err
		}
		if err := q.packData("注册成功"); err != nil {
			return nil, err
		}
	} else {
		if err := q.loginFlow(); err != nil {
			q.data = &response.UserRegisterResponse{
				Response: response.Response{
					StatusCode: 4,
					StatusMsg:  err.Error(),
				}}
			return q.data, err
		}
		if err := q.packData("登录成功"); err != nil {
			return nil, err
		}
	}
	return q.data, nil
}

// 对参数进行再次校验
func (q *LoginRegisterFlow) checkParams() error {
	if q.username == "" || q.password == "" {
		return errors.New("用户名或密码不能为空")
	}
	if len(q.username) > MaxUsernameLength {
		return errors.New("用户名长度须小于32")
	}
	if len(q.password) < MinPasswordLength || len(q.password) > MaxPasswordLength {
		return errors.New("密码长度必须大于8且小于32")
	}
	return nil
}

// 登录处理逻辑
func (q *LoginRegisterFlow) loginFlow() error {
	user, err := UserRegisterDao.FindUserByNameAndPass(q.username, q.password)
	if err != nil {
		global.Log.Error("UserRegisterService.LoginFlow USE UserRegisterDao.FindUserByNameAndPass ERROR ", zap.Error(err))
		return errors.New("用户名或密码错误")
	}
	// 颁发token
	token, err := jwt.GenToken(jwt.PayLoad{Username: (*user).Username, UserID: (*user).ID})
	if err != nil {
		global.Log.Error("UserRegisterService.LoginFlow USE jwt.GenToken ERROR ", zap.Error(err))
		return errors.New("Token生成失败")
	}

	q.userid = (*user).ID
	q.token = token
	return nil
}

// 注册处理逻辑
func (q *LoginRegisterFlow) registerFlow() error {
	if exist := UserRegisterDao.IsUserExistByUsername(q.username); exist {
		return errors.New("该用户名已存在")
	}
	// 对用户名和密码进行数据库查询验证
	user, err := UserRegisterDao.AddNewUser(q.username, q.password)
	if err != nil {
		global.Log.Error("UserRegisterService.LoginFlow USE UserRegisterDao.AddNewUser ERROR ", zap.Error(err))
		return errors.New("添加用户失败")
	}
	// 颁发token
	token, err := jwt.GenToken(jwt.PayLoad{Username: (*user).Username, UserID: (*user).ID})
	if err != nil {
		global.Log.Error("UserRegisterService.RegisterFlow USE jwt.GenToken ERROR ", zap.Error(err))
		return errors.New("Token生成失败")
	}

	q.userid = (*user).ID
	q.token = token
	return nil
}

// 对返回的数据进行打包
func (q *LoginRegisterFlow) packData(msg string) error {
	q.data = &response.UserRegisterResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  msg,
		},
		Token:  q.token,
		UserId: q.userid,
	}
	return nil
}
