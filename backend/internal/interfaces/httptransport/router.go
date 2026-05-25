package httptransport

import (
	"backend/internal/application/auth"
	cartoonuc "backend/internal/application/cartoon"

	"github.com/gofiber/fiber/v3"
)

// Deps 聚合注入到 HTTP 层的应用服务。
type Deps struct {
	Login        *auth.Login
	ListCartoons *cartoonuc.ListCartoons
}

// Register 注册路由（具体处理函数见 handlers_*.go）。
func Register(app *fiber.App, d Deps) {
	app.Get("/health", Health)

	app.Post("/login", func(c fiber.Ctx) error { return login(c, d.Login) })

	api := app.Group("/api/v1")
	api.Post("/auth/login", func(c fiber.Ctx) error { return login(c, d.Login) })
	api.Get("/cartoons", func(c fiber.Ctx) error { return listCartoons(c, d.ListCartoons) })
}
