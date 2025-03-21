package routers

import (
	"fmt"
	"discountmodule/configs"
	"discountmodule/controllers"
	"discountmodule/frameworks"
	"discountmodule/usecases"
)


func SetService(app *configs.AppConfig) {
	fmt.Println("setService")
	NewHttpReqest := frameworks.NewHttpReqest()
	NewServiceUsecase := usecases.NewServiceUsecase(NewHttpReqest)
	NewServiceController := controllers.NewServiceController(NewServiceUsecase)
	BaseGroup := app.App.Group(configs.App.Prefix)

	BaseGroup.Get("/discountmodule",
		NewServiceController.DisCountModule,
	)
}
