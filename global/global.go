package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// DB 数据库
var DB *gorm.DB

// RDB0 redis 0 client
var RDB0 *redis.Client

// Trans 全局翻译器
var Trans ut.Translator
