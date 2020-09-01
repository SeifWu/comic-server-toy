package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthSessionMiddleware Session 中间件
func AuthSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		fmt.Println(sessionValue)
		// 设置简单的变量
		c.Set("userId", sessionValue.(uint))
		c.Next()
		return
	}

}
