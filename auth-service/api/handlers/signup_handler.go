package handlers

import (
	"auth-service/domain"
	"auth-service/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func SignupHandler(registrationService service.Registration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var registrationUser domain.RegisterData
		if err := c.BodyParser(&registrationUser); err != nil {
			log.Println(err)
			return c.Status(http.StatusForbidden).JSON(err.Error())
		}

		loginResult, err := registrationService.SignUp(registrationUser)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusForbidden).JSON(err.Error())
		}

		return c.JSON(fiber.Map{
			"access_token":  loginResult.AccessToken,
			"refresh_token": loginResult.RefreshToken,
		})
	}
}
