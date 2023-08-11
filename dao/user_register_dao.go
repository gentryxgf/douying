package dao

import (
	"douyin/common/global"
	"douyin/models"
	"time"

	"go.uber.org/zap"
)

type UserRegisterDao struct{}

// 通过用户名称查询是否存在
func (UserRegisterDao) IsUserExistByUsername(username string) bool {
	var user models.UserModel
	global.DB.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

// 添加用户
func (UserRegisterDao) AddNewUser(name, password string) (*models.UserModel, error) {
	var user = models.UserModel{
		MODEL:    models.MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Username: name,
		Password: password}
	err := global.DB.Create(&user).Error
	if err != nil {
		global.Log.Error("添加用户失败", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

// 通过用户名和密码查询用户
func (UserRegisterDao) FindUserByNameAndPass(name, password string) (*models.UserModel, error) {
	var user models.UserModel
	err := global.DB.Where("username = ? and password = ?", name, password).First(&user).Error
	if err != nil {
		global.Log.Error("登录查询用户失败", zap.String("name", name), zap.Error(err))
		return nil, err
	}
	return &user, nil
}

// 通过ID查询用户
func (UserRegisterDao) FindUserById(id int64) (*models.UserModel, error) {
	var user models.UserModel
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		global.Log.Error("ID查询用户失败", zap.Int64("user_id", id), zap.Error(err))
		return nil, err
	}
	return &user, nil
}
