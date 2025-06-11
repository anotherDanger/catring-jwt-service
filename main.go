package main

import (
	"catering-jwt-service/app"
	"catering-jwt-service/controller"
	"catering-jwt-service/service"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	svc := service.NewServiceImpl()
	ctrl := controller.NewControllerImpl(svc)

	app := app.GetFiberApp()
	app.Use(cors.New())
	app.Post("/v1/auth", ctrl.Register)

	app.Listen(":8081")
}
