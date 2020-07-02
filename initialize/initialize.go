package initialize

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"seifwu.com/gin-basic-project/global"
)

// InitDB 初始化数据库
func InitDB() *gorm.DB {
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

	DBMigrate()

	global.DB = db
	return db
}
