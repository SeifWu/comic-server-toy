package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

// DB 数据库
var DB *gorm.DB

// RDB0 redis 0 client
var RDB0 *redis.Client
