package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
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
		})
		fmt.Println("Fiber new App")
	})
	return appInstance
}
