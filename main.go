package main

import (
	application "seifwu/config"
	initializers "seifwu/config/initializers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initializers.Initializers()
	application.Run()
}
