package routers

import (
	v1customerapi "seifwu/app/controllers/api/v1/customer"
	v1managerapi "seifwu/app/controllers/api/v1/manager"

	"github.com/gin-gonic/gin"
)

// UserRoutes 路由
func UserRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		customer := v1.Group("/customer/user")
		{
			customer.POST("", v1customerapi.Register)
			customer.DELETE("/:id", v1customerapi.UnsubscribeUser)
			customer.GET("/:id", v1customerapi.FindUser)
		}
		manager := v1.Group("/manager/user")
		{
			manager.GET("", v1managerapi.FindListUser)
		}
	}

}
