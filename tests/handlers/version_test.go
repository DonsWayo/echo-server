package handlers_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/a-safe-digital/echo-server/internal/config"
	"github.com/a-safe-digital/echo-server/internal/handlers"
	"github.com/a-safe-digital/echo-server/internal/platform/constants"
	"github.com/a-safe-digital/echo-server/internal/version"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestVersionEndpoint(t *testing.T) {
	app := fiber.New()
	cfg := config.NewDefaultConfig()
	handlers.RegisterRoutes(app, cfg)

	req := httptest.NewRequest(constants.MethodGET, constants.RouteVersion, nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, constants.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var versionInfo map[string]string
	err = json.Unmarshal(body, &versionInfo)
	assert.NoError(t, err)

	expectedInfo := version.GetBuildInfo()
	assert.Equal(t, expectedInfo["version"], versionInfo["version"])
	assert.Equal(t, expectedInfo["goVersion"], versionInfo["goVersion"])
}
