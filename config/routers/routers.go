package routers

import (
	v1api "seifwu/app/controllers/api/v1"

	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// 发送邮件
		v1.POST("/send_mail", v1api.SendAuthCodeMailsController)
		// 登录
		v1.POST("/login", v1api.Login)
	}

	UserRoutes(router)

	return router
}
