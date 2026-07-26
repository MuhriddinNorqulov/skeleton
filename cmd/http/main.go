package main

import (
	container "github.com/muhriddinnorqulov/skeleton/cmd/container"
)

// @title Skeleton API
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
