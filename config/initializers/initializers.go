package initializers

import (
	"fmt"
	"os"

	"seifwu/global"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Initializers 集合初始化
func Initializers() {
	InitConfig()
	InitDB()
	InitRedis()
}

// InitRedis redis 配置初始化
func InitRedis() {
	var c *gin.Context
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(c).Result()
	if err != nil {
		panic("redis ping error")
	}

	global.RDB0 = RedisClient
}

// InitConfig 配置初始化
func InitConfig() {
	rootedPath, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(rootedPath + "/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// InitDB 初始化数据库
func InitDB() {
	driverName := viper.GetString("dataSource.driverName")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		viper.GetString("dataSource.username"),
		viper.GetString("dataSource.password"),
		viper.GetString("dataSource.host"),
		viper.GetString("dataSource.port"),
		viper.GetString("dataSource.database"),
		viper.GetString("dataSource.charset"),
	)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("数据库连接失败，错误:" + err.Error())
	}

	global.DB = db
	DBMigrate()
}
