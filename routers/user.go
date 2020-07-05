package routers

import (
	"github.com/gin-gonic/gin"
	v1controller "seifwu.com/gin-basic-project/controller/api/v1"
)

// UserRoutes 路由
func UserRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/v1/user", v1controller.Register)
}
