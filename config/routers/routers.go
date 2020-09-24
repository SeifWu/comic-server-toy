package router

import (
	v1api "seifwu/app/controllers/api/v1"
	"seifwu/app/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Routers 路由
func Routers() *gin.Engine {
	router := gin.Default()
	sessionSecretKey := viper.GetString("session.secretKey")

	// store := cookie.NewStore([]byte(""))
	// router.Use(sessions.Sessions("appSession", store))

	store, _ := redis.NewStoreWithDB(10, "tcp", "localhost:6379", "", "1", []byte(sessionSecretKey))
	store.Options(sessions.Options{Path: "/", MaxAge: 2629746, HttpOnly: true})
	router.Use(sessions.Sessions("appSession", store))

	v1 := router.Group("/api/v1")
	{
		v1public := v1.Group("/public")
		V1Public(v1public)
		v1manager := v1.Group("/manager")
		v1manager.Use(middleware.JWTAuthMiddleware())
		V1Manager(v1manager)
		// 发送邮件
		v1.POST("/send_mail", v1api.SendAuthCodeMailsController)
	}

	UserRoutes(router)

	return router
}
