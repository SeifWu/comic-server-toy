package v1customerapi

import (
	"net/http"

	param "seifwu/app/params"
	service "seifwu/app/services"
	"seifwu/global/response"

	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(c *gin.Context) {
	var registerParams param.RegisterParams

	err := c.ShouldBindJSON(&registerParams)
	if err != nil {
		response.Fail(c, gin.H{"errMsg": "传递参数有误"})
		return
	}

	result := service.UserCreateService(c, registerParams)
	if result["success"] == false {
		response.Response(c, http.StatusBadRequest, 400, nil, result["message"])
		return
	}

	response.Success(c, gin.H{"user": result["data"]}, result["message"])
}

// UnsubscribeUser 注销用户
func UnsubscribeUser(c *gin.Context) {
	// TODO 流程
	// 退出并删除用户
	id := c.Param("id")

	result := service.UserUnsubscribeService(c, id)
	if result["success"] == false {
		response.Response(c, http.StatusBadRequest, 400, nil, result["message"])
		return
	}

	response.Response(c, http.StatusOK, 200, nil, "注销成功")
}

// FindUser GET /api/v1/user/:id
func FindUser(c *gin.Context) {
	id := c.Param("id")

	result := service.UserFindService(c, id)
	if result["success"] == false {
		response.Response(c, http.StatusBadRequest, 400, nil, result["message"])
		return
	}

	response.Response(c, http.StatusOK, 200, gin.H{"user": result["data"]}, result["message"])
}
