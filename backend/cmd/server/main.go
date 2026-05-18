package main

import (
	"backend/internal/db"
	"backend/internal/handler"

	"github.com/gofiber/fiber/v3"
)

func main() {
	db.InitMySQL()

	app := fiber.New()

	app.Post("/login", handler.Login)

	app.Listen(":8080")
}
