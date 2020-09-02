package v1api

import (
	"log"
	"net/http"
	model "seifwu/app/models"
	param "seifwu/app/params"
	global "seifwu/global"
	"seifwu/global/response"
	util "seifwu/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login 定义登陆逻辑
// model.LoginReq中定义了登陆的请求体(userName,passWord)
func Login(c *gin.Context) {
	DB := global.DB

	var loginParams param.LoginParams
	var user model.User

	err := c.ShouldBindJSON(&loginParams)
	if err != nil {
		response.Fail(c, gin.H{"errMsg": "传递参数有误"})
		return
	}

	// 登陆逻辑校验(查库，验证用户是否存在以及登陆信息是否正确)
	// 判断手机号是否存在
	DB = DB.Where(&model.User{UserName: loginParams.UserName}).First(&user)

	if user.ID == 0 {
		response.Response(c, http.StatusBadRequest, 40001, nil, "用户不存在")
		return
	}

	// 密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParams.PassWord)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	// 验证通过后为该次请求生成token
	if user.ID != 0 {
		generateToken(c, user)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": -1, "msg": "验证失败", "data": user})
	}
}

// token生成器
// md 为上面定义好的middleware中间件
func generateToken(c *gin.Context, user model.User) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := util.NewJWT()
	// 构造用户claims信息(负荷)
	claims := util.CustomClaims{
		user.UserName,
		user.Email,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "Seifwu",                        // 签名颁发者
		},
	} // 根据claims生成token对象
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": -1, "msg": err.Error(), "data": nil})
	}
	log.Println(token)
	// 封装一个响应数据,返回用户名和token
	data := gin.H{"userName": user.UserName, "token": token}
	c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "登陆成功", "data": data})
	return
}
