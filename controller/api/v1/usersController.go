package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"seifwu.com/gin-basic-project/global"
	"seifwu.com/gin-basic-project/global/response"
	"seifwu.com/gin-basic-project/model"
	"seifwu.com/gin-basic-project/utils"
)

// RegisterParams 组册参数
type RegisterParams struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
}

// Register 注册
func Register(c *gin.Context) {
	var registerParams RegisterParams
	DB := global.DB

	err := c.BindJSON(&registerParams)
	if err != nil {
		log.Fatal(err)
	}

	userName := registerParams.Username
	password := registerParams.Password

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于 6 位")
		return
	}

	// 如果名称不存在，赋值随机字符串
	if len(userName) == 0 {
		userName = utils.RandomString(10)
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "加密错误")
		return
	}

	newUser := model.User{
		UserName: userName,
		Password: string(hasedPassword),
	}

	DB.Create(&newUser)

	// 返回结果
	response.Success(c, nil, "注册成功")
}
