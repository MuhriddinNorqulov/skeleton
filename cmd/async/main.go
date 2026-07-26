package main

import container "github.com/muhriddinnorqulov/skeleton/cmd/container"

func main() {
	app := container.InitAsyncApp()
	app.Init()
	app.Start()
}
