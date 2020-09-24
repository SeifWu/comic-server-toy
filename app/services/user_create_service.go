package service

import (
	model "seifwu/app/models"
	param "seifwu/app/params"
	"seifwu/global"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

// UserCreateService 创建用户服务
func UserCreateService(c *gin.Context, params param.SignUpParam) gin.H {
	DB := global.DB
	var user model.User
	userName := params.Username
	password := params.Password
	authCode := params.AuthCode
	email := params.Email

	if len(authCode) == 0 {
		return gin.H{"success": false, "message": "验证码不能为空"}
	}

	// 暂时只有 Email 注册
	getAuthCode := global.RDB0.Get(c, params.Email+"-register")
	if err := getAuthCode.Err(); err != nil {
		if err == redis.Nil {
			return gin.H{"success": false, "message": "验证码过期或不存在"}
		}
		panic(err)
	}

	if getAuthCode.Val() != authCode {
		return gin.H{"success": false, "message": "验证码不正确"}
	}

	if len(password) < 6 {
		return gin.H{"success": false, "message": "密码不能少于 6 位"}
	}

	if len(userName) == 0 {
		return gin.H{"success": false, "message": "用户名不能为空"}
	}

	var count int
	DB.Where("user_name = ? ", userName).First(&user).Count(&count)
	if count > 0 {
		return gin.H{"success": false, "message": "该用户名已存在"}
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return gin.H{"success": false, "message": "加密错误"}
	}

	newUser := model.User{
		Username: userName,
		NickName: userName,
		Email:    email,
		Password: string(hasedPassword),
	}

	if createUser := DB.Create(&newUser); createUser.Error != nil {
		return gin.H{"success": false, "message": createUser.Error}
	}

	// 返回结果
	return gin.H{"success": true, "message": "注册成功", "data": newUser}
}
