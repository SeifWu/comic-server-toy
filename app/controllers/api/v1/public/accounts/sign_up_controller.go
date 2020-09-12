package v1customeraccountapi

import (
	"net/http"
	param "seifwu/app/params"
	service "seifwu/app/services"
	"seifwu/global/response"

	"github.com/gin-gonic/gin"
)

// SignUp 注册
func SignUp(c *gin.Context) {
	var signUpParam param.SignUpParam

	// TODO 系统化参数验证
	if err := c.ShouldBindJSON(&signUpParam); err != nil {
		response.Fail(c, gin.H{"errMsg": "传递参数有误"})
		return
	}

	result := service.UserCreateService(c, signUpParam)
	if result["success"] == false {
		response.Response(c, http.StatusBadRequest, 400, nil, result["message"])
		return
	}

	response.Success(c, gin.H{"user": result["data"]}, result["message"])
}
