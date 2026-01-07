package handlers

import (
	"api-gateway/utils"

	"github.com/gofiber/fiber/v2"
)

func GuestToken(c *fiber.Ctx) error {
	token, err := utils.GenerateGuestToken()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"access_token": token,
		"expires_in":   600,
	})
}
