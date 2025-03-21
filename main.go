package main

import (
	"discountmodule/configs"
	"discountmodule/routers"
)

func main() {
	app := configs.NewApp()
	app.SetApp()

	routers.SetService(app)
	app.RunApp()
}
