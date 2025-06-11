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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-type, Accept, Authorization",
		AllowMethods:     "POST",
	}))
	app.Post("/v1/auth", ctrl.Register)
	app.Post("/v1/refresh", ctrl.Refresh)
	app.Listen(":8081")
}
