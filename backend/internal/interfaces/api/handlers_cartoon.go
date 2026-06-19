package api

import (
	cartoonuc "backend/internal/application/cartoon"

	"github.com/gofiber/fiber/v3"
)

func listCartoons(c fiber.Ctx, uc *cartoonuc.ListCartoons) error {
	items, err := uc.Execute(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "list failed"})
	}
	return c.JSON(fiber.Map{"data": items})
}
