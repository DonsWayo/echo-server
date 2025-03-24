package handlers

import (
	"os"

	"github.com/a-safe-digital/echo-server/internal/config"
	"github.com/a-safe-digital/echo-server/internal/platform/constants"
	"github.com/a-safe-digital/echo-server/internal/platform/utils"
	"github.com/a-safe-digital/echo-server/internal/version"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App, config *config.Configuration) {
	app.Get(constants.RouteHealth, func(c fiber.Ctx) error {
		return c.SendString(constants.MessageHealthOK)
	})

	app.Get("/version", func(c fiber.Ctx) error {
		return c.JSON(version.GetBuildInfo())
	})

	app.Get(constants.RouteRoot, func(c fiber.Ctx) error {
		if filePath := utils.GetValueFromHeaderOrQuery(c, config.Commands.File.Header, config.Commands.File.Query); filePath != "" && config.EnableFile {
			return handleFileExploration(c, filePath)
		}
		if envVar := utils.GetValueFromHeaderOrQuery(c, config.Commands.HTTPEnvBody.Header, config.Commands.HTTPEnvBody.Query); envVar != "" && config.EnableEnvironment {
			return c.JSON(os.Getenv(envVar))
		}
		if body := utils.GetValueFromHeaderOrQuery(c, config.Commands.HTTPBody.Header, config.Commands.HTTPBody.Query); body != "" {
			return c.JSON(body)
		}
		response := buildDefaultResponse(c, config)
		return c.JSON(response)
	})

	app.All(constants.RouteRoot, func(c fiber.Ctx) error {
		body := c.Body()
		if len(body) > 0 {
			return c.Send(body)
		}
		if filePath := utils.GetValueFromHeaderOrQuery(c, config.Commands.File.Header, config.Commands.File.Query); filePath != "" && config.EnableFile {
			return handleFileExploration(c, filePath)
		}
		if envVar := utils.GetValueFromHeaderOrQuery(c, config.Commands.HTTPEnvBody.Header, config.Commands.HTTPEnvBody.Query); envVar != "" && config.EnableEnvironment {
			return c.JSON(os.Getenv(envVar))
		}
		if body := utils.GetValueFromHeaderOrQuery(c, config.Commands.HTTPBody.Header, config.Commands.HTTPBody.Query); body != "" {
			return c.JSON(body)
		}
		response := buildDefaultResponse(c, config)
		return c.JSON(response)
	})
}
