package middleware

import (
	"net/http"
	util "seifwu/utils"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// // 按空格分割
		// parts := strings.SplitN(authHeader, " ", 2)
		// if !(len(parts) == 2 && parts[0] == "Bearer") {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"code": 2004,
		// 		"msg":  "请求头中auth格式有误",
		// 	})
		// 	c.Abort()
		// 	return
		// }
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := util.ParseJWT(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
