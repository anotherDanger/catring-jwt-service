package main

import (
	"catering-jwt-service/controller"
	"catering-jwt-service/service"
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	appInstance *fiber.App
	once        sync.Once
)

func GetFiberApp() *fiber.App {
	once.Do(func() {
		appInstance = fiber.New(fiber.Config{
			AppName:      "JWT Service",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  10 * time.Second,
			ProxyHeader:  "X-Forwarded-Proto",
		})
		fmt.Println("Fiber new App")
	})
	return appInstance
}

func main() {
	svc := service.NewServiceImpl()
	ctrl := controller.NewControllerImpl(svc)

	app := GetFiberApp()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://khatering.shop, http://localhost:3000",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-type, Accept, Authorization",
		AllowMethods:     "POST",
	}))
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/v1/auth", ctrl.Register)
	app.Post("/v1/refresh", ctrl.Refresh)
	app.Post("/v1/logout", ctrl.LogoutHandler)
	app.Listen(":8081")
}
