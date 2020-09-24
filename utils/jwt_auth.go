package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// JwtExpireDuration JWT 过期时间
const JwtExpireDuration = time.Hour * 2

// JwtSecretKey 生成 Token 时的密钥
var JwtSecretKey = []byte(viper.GetString("jwt.secretKey"))

// JWTAuthClaims 自定义声明结构体并内嵌jwt.StandardClaims
type JWTAuthClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT 生成 Token
func GenerateJWT(username string) (string, error) {
	jwtAuthClaims := JWTAuthClaims{
		"username", // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JwtExpireDuration).Unix(), // 过期时间
			Issuer:    viper.GetString("jwt.issuer"),            // 签发人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtAuthClaims)
	token, err := tokenClaims.SignedString(JwtSecretKey)

	return token, err
}

// ParseJWT 解析 JWT
func ParseJWT(token string) (*JWTAuthClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JWTAuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JWTAuthClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
