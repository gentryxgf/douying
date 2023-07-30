package service

import (
	"douying/middleware"
	"douying/model"
	"errors"
)

const (
	MaxUsernameLength = 32
)

type LoginResponse struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

type LoginRegisterFlow struct {
	username string         //前端传回的username
	password string         //前端传回的password
	data     *LoginResponse //真正返回的数据
	userid   uint           //用于返回的userid
	token    string         //用于返回的token
}

func UserRegister(username, password string) (*LoginResponse, error) {
	return NewLoginRegisterFlow(username, password).Do("register")
}

func UserLogin(username, password string) (*LoginResponse, error) {
	return NewLoginRegisterFlow(username, password).Do("login")
}

func NewLoginRegisterFlow(username, password string) *LoginRegisterFlow {
	return &LoginRegisterFlow{username: username, password: password}
}

func (q *LoginRegisterFlow) Do(method string) (*LoginResponse, error) {
	if err := q.checkParams(); err != nil {
		return nil, err
	}

	if method == "login" {
		if err := q.loginFlow(); err != nil {
			return nil, err
		}
	} else {
		if err := q.registerFlow(); err != nil {
			return nil, err
		}
	}

	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.data, nil
}

// 对参数进行校验
func (q *LoginRegisterFlow) checkParams() error {
	if q.username == "" {
		return errors.New("用户名为空")
	}
	if len(q.username) > MaxUsernameLength {
		return errors.New("用户名长度超过限制")
	}
	if q.password == "" {
		return errors.New("密码不正确。密码长度必须大于8且小于32")
	}
	return nil
}

// 登录流程
func (q *LoginRegisterFlow) loginFlow() error {
	//判断是否在登录用户缓存表内
	if userCache, ok := model.UserLoginInfo[q.username]; ok {
		// 判断缓存内的密码与登录密码是否相同
		if userCache.Password != q.password {
			return errors.New("密码不正确")
		} else {
			// 用户在缓存内且密码正确，重新颁发token返回
			token, _ := middleware.GenerateToken(userCache.ID)
			q.userid = userCache.ID
			q.token = token
			return nil
		}
	} else {
		// 用户不在登录用户缓存表内，则通过用户名和密码去数据库查找是否有该用户
		user, err := model.NewUserDaoInstance().QueryUserByNameAndPassword(q.username, q.password)
		if err != nil {
			return errors.New("该用户不存在")
		}
		// 用户在数据库内，则颁发token并返回，同时将用户加入登录用户缓存表
		token, _ := middleware.GenerateToken(user.ID)
		q.userid = (*user).ID
		q.token = token
		model.UserLoginInfo[q.username] = *user
		return nil
	}
}

// 注册流程
func (q *LoginRegisterFlow) registerFlow() error {
	db := model.NewUserDaoInstance()
	// 校验用户名是否存在
	_, err := db.QueryUserByName(q.username)
	if err == nil {
		return errors.New("该用户名已存在")
	}
	// 将用户添加进数据库
	user, err := db.AddUser(q.username, q.password)
	if err != nil {
		return errors.New("用户添加失败")
	}
	// 颁发token以及将用户加入登录缓存表
	token, _ := middleware.GenerateToken((*user).ID)
	q.userid = (*user).ID
	q.token = token
	model.UserLoginInfo[q.username] = *user
	return nil
}

// 构造返回的数据
func (q *LoginRegisterFlow) packData() error {
	q.data = &LoginResponse{
		UserId: q.userid,
		Token:  q.token,
	}
	return nil
}
