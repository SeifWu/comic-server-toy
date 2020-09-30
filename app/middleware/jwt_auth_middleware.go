package middleware

import (
	"context"
	"net/http"
	"seifwu/global"
	"seifwu/global/response"
	util "seifwu/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Response(c, http.StatusUnauthorized, "401", nil, "请先登录", nil)
			c.Abort()
			return
		}

		mc, err := util.ParseJWT(authHeader)
		tokenString := ""
		if err != nil {
			validationError, hasError := err.(*jwt.ValidationError)
			errorMessage := "Token 错误"
			if hasError {
				if validationError.Errors&jwt.ValidationErrorExpired != 0 {
					errorMessage = "Token 过期"
					val, err := global.RDB0.Get(ctx, mc.UUID).Result()
					// 尝试刷新 Token
					// Token 是过期的，Redis 有缓存且两值相等
					if validationError.Errors&jwt.ValidationErrorExpired != 0 && err == nil && authHeader == val {
						tokenString, _ = util.GenerateJWT(mc.Username)
					}
				}
				// TODO 添加其他错误
			}

			if tokenString != "" {
				c.Writer.Header().Add("X-Token", tokenString)
				c.Set("username", mc.Username)
				c.Next()
			} else {
				response.Response(c, http.StatusUnauthorized, "401", nil, errorMessage, nil)
				c.Abort()
				return
			}
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
