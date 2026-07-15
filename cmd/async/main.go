package main

import container "example.com/PROJECT_NAME/cmd/container"

func main() {
	app := container.InitAsyncApp()
	app.Init()
	app.Start()
}
