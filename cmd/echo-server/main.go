package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/a-safe-digital/echo-server/internal/config"
	"github.com/a-safe-digital/echo-server/internal/handlers"
	"github.com/a-safe-digital/echo-server/internal/middleware"
	"github.com/a-safe-digital/echo-server/internal/platform/logger"
)

func main() {
	cfg := config.NewDefaultConfig()
	app := fiber.New()

	// Print server information
	logger.PrintServerInfo(cfg)

	// Register middleware
	app.Use(logger.Logger(cfg))
	app.Use(middleware.ProcessRequest(cfg))

	// Register routes
	handlers.RegisterRoutes(app, cfg)

	log.Fatal(app.Listen(":" + cfg.Port))
}
