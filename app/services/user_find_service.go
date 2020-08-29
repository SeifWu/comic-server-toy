package service

import (
	model "seifwu/app/models"
	"seifwu/global"

	"github.com/gin-gonic/gin"
)

// UserFindService 注销用户服务
func UserFindService(c *gin.Context, id string) gin.H {
	DB := global.DB
	var user model.User

	DB.First(&user, id)
	if user.ID == 0 {
		return gin.H{"success": false, "message": "该用户名不存在"}
	}

	return gin.H{"success": true, "data": user, "message": "获取成功"}
}
