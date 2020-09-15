package router

import (
	v1public "seifwu/app/controllers/api/v1/public"

	"github.com/gin-gonic/gin"
)

// V1Public v1 public 接口
func V1Public(router *gin.RouterGroup) {
	router.POST("/sign_up", v1public.SignUp)
}
