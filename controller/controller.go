package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Register(c *fiber.Ctx) error
}
