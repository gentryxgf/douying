package model

import (
	"log"
	"sync"

	"gorm.io/gorm"
)

// 登录用户缓存表，key为用户名，value为用户信息
var UserLoginInfo map[string]User

type User struct {
	gorm.Model
	Name     string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

// 初始化用户登录表
func (*UserDao) TokenMap() {
	UserLoginInfo = make(map[string]User, 50)
}

// 通过用户ID查询用户
func (*UserDao) QueryUserById(id uint) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Println("Find user by id failed: ", err.Error())
		return nil, err
	}
	return &user, nil
}

// 通过用户名查询用户
func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	err := db.Where("username = ?", name).First(&user).Error
	if err != nil {
		log.Println("Find user by name failed: ", err.Error())
		return nil, err
	}
	return &user, nil
}

// 通过用户名和密码查询用户
func (*UserDao) QueryUserByNameAndPassword(name, password string) (*User, error) {
	var user User
	err := db.Where("username = ? AND password = ?", name, password).First(&user).Error
	if err != nil {
		log.Println("Find user by name and password failed: ", err.Error())
		return nil, err
	}
	return &user, nil
}

// 添加用户
func (*UserDao) AddUser(name, password string) (*User, error) {
	user := User{Name: name, Password: password}
	err := db.Create(&user).Error
	if err != nil {
		log.Println("Add user failed: ", err.Error())
		return nil, err
	}
	return &user, nil
}
