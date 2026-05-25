package main

import (
	"backend/internal/application/auth"
	cartoonuc "backend/internal/application/cartoon"
	"backend/internal/config"
	infradb "backend/internal/infrastructure/db"
	"backend/internal/infrastructure/persistence/mysqlstore"
	"backend/internal/interfaces/httptransport"
	"backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := config.Load()

	gormDB, err := infradb.Open(cfg.MySQLDSN)
	if err != nil {
		panic(err)
	}

	userRepo := mysqlstore.NewUserRepository(gormDB)
	cartoonRepo := mysqlstore.NewCartoonRepository(gormDB)

	deps := httptransport.Deps{
		Login:        auth.NewLogin(userRepo),
		ListCartoons: cartoonuc.NewListCartoons(cartoonRepo),
	}

	app := fiber.New()
	middleware.Register(app)
	httptransport.Register(app, deps)

	_ = app.Listen(cfg.ListenAddr())
}
