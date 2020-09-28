package main

import (
	application "seifwu/config"
	initializers "seifwu/config/initializers"

	_ "gorm.io/driver/mysql"
)

func main() {
	initializers.Initializers()
	application.Run()
}
