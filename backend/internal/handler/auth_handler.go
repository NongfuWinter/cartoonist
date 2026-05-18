package handler

import (
	"backend/internal/service"

	"github.com/gofiber/fiber/v3"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	ok, err := service.Login(req.Username, req.Password)
	if err != nil || !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "login failed",
		})
	}

	return c.JSON(fiber.Map{
		"message": "login success",
	})
}
