package service

import (
	model "seifwu/app/models"
	"seifwu/global"

	"github.com/gin-gonic/gin"
)

// UserUnsubscribeService 注销用户服务
func UserUnsubscribeService(c *gin.Context, id string) gin.H {
	DB := global.DB
	var user model.User

	DB.First(&user, id).Unscoped().Delete(&user)
	return gin.H{"success": true, "message": "注销成功"}
}
