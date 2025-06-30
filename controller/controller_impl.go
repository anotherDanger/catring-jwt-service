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
	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	token, err := ctrl.svc.Register(c.Context(), &reqBody)
	if err != nil {
		return err
	}

	rToken, err := ctrl.svc.RefreshToken(c.Context(), &reqBody)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh",
		Value:    rToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})

	return c.JSON(web.Response{AccessToken: token, Username: reqBody.Username})
}

func (ctrl *ControllerImpl) Refresh(c *fiber.Ctx) error {

	refreshToken := c.Cookies("refresh")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "No refresh token 1"})
	}

	newAccessToken, username, err := ctrl.svc.Refresh(c.Context(), refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "No refresh token 2"})
	}

	return c.JSON(&web.Response{AccessToken: newAccessToken, Username: username})
}

func (ctrl *ControllerImpl) LogoutHandler(c *fiber.Ctx) error {

	c.Cookie(&fiber.Cookie{
		Name:     "refresh",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout berhasil",
	})
}
