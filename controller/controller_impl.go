package controller

import (
	"catering-jwt-service/domain"
	"catering-jwt-service/service"
	"catering-jwt-service/web"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ControllerImpl struct {
	svc service.Service
}

func NewControllerImpl(svc service.Service) Controller {
	return &ControllerImpl{
		svc: svc,
	}
}

func (ctrl *ControllerImpl) Register(c *fiber.Ctx) error {
	var reqBody domain.Domain
	c.BodyParser(&reqBody)
	token, err := ctrl.svc.Register(c.Context(), &reqBody)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh",
		Expires:  time.Now().Add(7 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	})

	c.JSON(web.Response{AccessToken: token})
	return nil
}
