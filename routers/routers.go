package routers

import (
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers() *gin.Engine {
	router := gin.Default()
	UserRoutes(router)

	return router
}
