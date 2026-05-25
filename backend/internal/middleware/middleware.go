package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

// Register attaches global middleware to the app.
func Register(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New())
}
