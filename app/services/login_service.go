package service

import (
	"fmt"
	model "seifwu/app/models"
	v1serializer "seifwu/app/serializers/api/v1/manager"
	"seifwu/global"

	"golang.org/x/crypto/bcrypt"
)

// LoginService 用户注册服务
func LoginService(user *model.User) (resultUser *v1serializer.UserSerializer, err error) {

	DB := global.DB
	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hasedPassword)

	if err != nil {
		return nil, fmt.Errorf("密码加密错误")
	}

	if createUser := DB.Create(user).Scan(&newUser); createUser.Error != nil {
		return nil, createUser.Error
	}

	return &newUser, nil
}
