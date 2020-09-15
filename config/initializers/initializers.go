package initializers

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"seifwu/global"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Initializers 集合初始化
func Initializers() {
	InitConfig()
	InitDB()
	InitRedis()
	if err := InitTrans("zh"); err != nil {
		panic(err)
	}
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

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, zhT)

		var ok bool
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			log.Println("local: en")
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			log.Println("local: zh")
			err = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			log.Println("local: zh")
			err = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}
