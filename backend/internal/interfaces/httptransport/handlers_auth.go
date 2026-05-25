package httptransport

import (
	"backend/internal/application/auth"

	"github.com/gofiber/fiber/v3"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c fiber.Ctx, uc *auth.Login) error {
	var req loginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	ok, err := uc.Execute(c.Context(), req.Username, req.Password)
	if err != nil || !ok {
		return c.Status(401).JSON(fiber.Map{"error": "login failed"})
	}
	return c.JSON(fiber.Map{"message": "login success"})
}
