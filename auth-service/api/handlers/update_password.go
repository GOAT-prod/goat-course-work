package handlers

import (
	"auth-service/domain"
	"auth-service/service"
	"github.com/gofiber/fiber/v2"
)

func UpdatePasswordHandler(authService service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var updatePassword domain.UpdatePasswordRequest
		if err := c.BodyParser(&updatePassword); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if err := authService.UpdatePassword(updatePassword); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	}
}
