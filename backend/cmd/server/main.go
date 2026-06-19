package main

import (
	"backend/internal/application/auth"
	cartoonuc "backend/internal/application/cartoon"
	"backend/internal/config"
	infradb "backend/internal/infrastructure/db"
	"backend/internal/infrastructure/persistence/mysqlstore"
	"backend/internal/interfaces/api"
	"backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// load config
	cfg := config.Load()

	// open database
	gormDB, err := infradb.Open(cfg.MySQLDSN)
	if err != nil {
		panic(err)
	}

	// create repositories
	userRepo := mysqlstore.NewUserRepository(gormDB)
	cartoonRepo := mysqlstore.NewCartoonRepository(gormDB)

	// create dependencies
	deps := api.Deps{
		Login:        auth.NewLogin(userRepo),
		ListCartoons: cartoonuc.NewListCartoons(cartoonRepo),
	}

	// create fiber app
	app := fiber.New()

	// register middleware
	middleware.Register(app)

	// register api
	api.Register(app, deps)

	// start server
	_ = app.Listen(cfg.ListenAddr())
}
