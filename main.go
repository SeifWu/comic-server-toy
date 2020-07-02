package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"seifwu.com/gin-basic-project/config"
	"seifwu.com/gin-basic-project/global"
	"seifwu.com/gin-basic-project/initialize"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitConfig()
	initialize.InitDB()
	defer global.DB.Close()

	r := gin.Default()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 默认监听并在 0.0.0.0:8080 上启动服务
}
