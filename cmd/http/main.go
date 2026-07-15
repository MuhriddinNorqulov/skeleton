package main

import (
	container "example.com/PROJECT_NAME/cmd/container"
)

// @title PROJECT_NAME API
// @version 1.0
// @BasePath /api/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @securityDefinitions.basic BasicAuth
func main() {
	app := container.InitHttpApp()
	app.Init()
	app.Start()
}
