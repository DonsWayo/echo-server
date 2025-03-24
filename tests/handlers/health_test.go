package handlers_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/a-safe-digital/echo-server/internal/config"
	"github.com/a-safe-digital/echo-server/internal/handlers"
	"github.com/a-safe-digital/echo-server/internal/platform/constants"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	app := fiber.New()
	cfg := config.NewDefaultConfig()
	handlers.RegisterRoutes(app, cfg)

	req := httptest.NewRequest(constants.MethodGET, constants.RouteHealth, nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, constants.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, constants.MessageHealthOK, string(body))
}
