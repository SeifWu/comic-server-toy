package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

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
