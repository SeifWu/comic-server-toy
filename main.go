package main

import (
	application "seifwu/config"
	initializers "seifwu/config/initializers"
)

func main() {
	initializers.Initializers()
	application.Run()
}
