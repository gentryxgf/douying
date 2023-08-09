package service

import (
	"douyin/common/jwt"
	"douyin/models/response"
	"errors"
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

func (UserRegisterService) UserRegister(username, password string) (*response.UserRegisterResponse, error) {
	return NewLoginRegisterFlow(username, password).Do("register")
}

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

func (q *LoginRegisterFlow) loginFlow() error {
	user, err := UserRegisterDao.FindUserByNameAndPass(q.username, q.password)
	if err != nil {
		return errors.New("用户名或密码错误")
	}

	token, err := jwt.GenToken(jwt.PayLoad{Username: (*user).Username, UserID: (*user).ID})
	if err != nil {
		return errors.New("Token生成失败")
	}

	q.userid = (*user).ID
	q.token = token
	return nil
}

func (q *LoginRegisterFlow) registerFlow() error {
	if exist := UserRegisterDao.IsUserExistByUsername(q.username); exist {
		return errors.New("该用户名已存在")
	}

	user, err := UserRegisterDao.AddNewUser(q.username, q.password)
	if err != nil {
		return errors.New("添加用户失败")
	}

	token, err := jwt.GenToken(jwt.PayLoad{Username: (*user).Username, UserID: (*user).ID})
	if err != nil {
		return errors.New("Token生成失败")
	}

	q.userid = (*user).ID
	q.token = token
	return nil
}

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
