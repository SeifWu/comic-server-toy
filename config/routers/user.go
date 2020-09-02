package routers

import (
	v1customerapi "seifwu/app/controllers/api/v1/customer"
	v1managerapi "seifwu/app/controllers/api/v1/manager"
	"seifwu/app/middleware"

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
		manager.Use(middleware.AuthSessionMiddleware())
		// manager.Use(csrf.Middleware(csrf.Options{
		// 	Secret: "iIsInR5cCI6IkpX9.eyJ1c2VyTmFtZSI6IkRv",
		// 	ErrorFunc: func(c *gin.Context) {
		// 		c.String(400, "CSRF token mismatch")
		// 		c.Abort()
		// 	},
		// }))
		{
			manager.GET("", v1managerapi.FindListUser)
		}
	}

}
