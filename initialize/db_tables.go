package initialize

import (
	"seifwu.com/gin-basic-project/global"
	"seifwu.com/gin-basic-project/model"
)

// DBTables Migrate Table
func DBTables() {
	db := global.DB

	db.AutoMigrate(
		&model.User{},
	)
}
