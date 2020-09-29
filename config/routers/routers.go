package router

import (
	v1api "seifwu/app/controllers/api/v1"
	v1manager "seifwu/app/controllers/api/v1/manager"

	"seifwu/app/middleware"

	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1public := v1.Group("/public")
		V1Public(v1public)

		// 管理端登录不需要鉴权
		v1.POST("/manager/login", v1manager.Login)
		v1manager := v1.Group("/manager")

		v1manager.Use(middleware.JWTAuthMiddleware())
		V1Manager(v1manager)
		// 发送邮件
		v1.POST("/send_mail", v1api.SendAuthCodeMailsController)
	}

	return router
}
