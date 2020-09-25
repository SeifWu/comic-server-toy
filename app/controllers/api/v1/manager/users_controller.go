package v1manager

import (
	"log"
	"net/http"

	model "seifwu/app/models"
	v1serializer "seifwu/app/serializers/api/v1/manager"
	service "seifwu/app/services"
	global "seifwu/global"
	"seifwu/global/response"

	"github.com/gin-gonic/gin"
)

// FindListUser GET /api/v1/manager/user
func FindListUser(c *gin.Context) {
	result := service.UserFindListService(c, true)
	if result["success"] == false {
		response.Fail(c, result)
		return
	}

	c.JSON(
		http.StatusOK,
		result,
	)
}

// CurrentUserController 获取当前用户
func CurrentUserController(c *gin.Context) {
	var userSerializer v1serializer.UserSerializer
	DB := global.DB
	username, isExists := c.Get("username")
	//
	var user model.User

	// DB = DB.Table("users").Scopes(scope.UserFindByUsername(username.(string))).Scan(&userSerializer)
	DB.Table("users").Where("username = ?", username.(string)).First(&user).Scan(&userSerializer)
	log.Println(username.(string), isExists, user)
	if !isExists || user.ID == 0 {
		response.Response(c, http.StatusUnauthorized, "40002", nil, "用户不存在", nil)
		return
	}

	response.Response(c, http.StatusOK, "0", gin.H{"user": userSerializer}, "请求成功", nil)
}
