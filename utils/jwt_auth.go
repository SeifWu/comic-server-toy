package util

import (
	"context"
	"seifwu/global"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var ctx = context.Background()

// JwtExpireDuration JWT 过期时间
const JwtExpireDuration = time.Hour * 2

// JwtSecretKey 生成 Token 时的密钥
var JwtSecretKey = []byte(viper.GetString("jwt.secretKey"))

// JWTAuthClaims 自定义声明结构体并内嵌jwt.StandardClaims
type JWTAuthClaims struct {
	Username string `json:"username"`
	UUID     string `json:"uuid"`
	jwt.StandardClaims
}

// GenerateJWT 生成 Token
func GenerateJWT(username string) (string, error) {
	uuid := uuid.Must(uuid.NewRandom()).String()
	jwtAuthClaims := JWTAuthClaims{
		username, // 自定义字段
		uuid,     // UUID
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JwtExpireDuration).Unix(), // 过期时间
			Issuer:    viper.GetString("jwt.issuer"),            // 签发人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtAuthClaims)
	token, err := tokenClaims.SignedString(JwtSecretKey)

	if err == nil {
		// 将 Token 设置到 Redis
		global.RDB0.Set(ctx, uuid, token, time.Hour*24*30)
	}

	return token, err
}

// ParseJWT 解析 JWT
func ParseJWT(token string) (*JWTAuthClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JWTAuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	claims, ok := tokenClaims.Claims.(*JWTAuthClaims)
	if ok && tokenClaims.Valid {
		return claims, nil
	}

	return claims, err
}
