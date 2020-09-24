package v1public

import (
	"net/http"
	model "seifwu/app/models"
	param "seifwu/app/params"
	global "seifwu/global"
	"seifwu/global/response"
	scope "seifwu/global/scopes"
	util "seifwu/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var loginParams param.LoginParams

// Login 定义登录逻辑
func Login(c *gin.Context) {
	// 参数错误处理
	if err := c.ShouldBind(&loginParams); err != nil {
		errResult := util.UnifiedValidation(c, err, "40001", loginParams)

		response.Response(c, http.StatusBadRequest, "40001", nil, errResult, nil)
		return
	}

	DB := global.DB

	var user model.User

	DB = DB.Scopes(scope.UserFindByUsername(loginParams.UserName)).First(&user)

	if user.ID == 0 {
		response.Response(c, http.StatusBadRequest, "40001", nil, "用户不存在", nil)
		return
	}

	// 密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParams.PassWord)); err != nil {
		response.Response(c, http.StatusBadRequest, "40001", nil, "密码错误", nil)
		return
	}

	if user.ID != 0 {
		tokenString, _ := util.GenerateJWT(user.Username)
		c.Writer.Header().Add("X-Token", tokenString)
		response.Response(c, http.StatusOK, "0", gin.H{"user": user}, nil, nil)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "msg": "验证失败", "data": user})
	}
}
