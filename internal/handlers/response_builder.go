package handlers

import (
	"os"
	"strings"

	"github.com/a-safe-digital/echo-server/internal/config"
	"github.com/gofiber/fiber/v3"
)

func buildDefaultResponse(c fiber.Ctx, config *config.Configuration) map[string]interface{} {
	response := make(map[string]interface{})
	
	if config.EnableHost {
		hostname, _ := os.Hostname()
		response["host"] = map[string]interface{}{
			"hostname": hostname,
			"pid":      os.Getpid(),
		}
	}

	if config.EnableHTTP {
		response["http"] = map[string]interface{}{
			"method":      c.Method(),
			"baseUrl":     c.BaseURL(),
			"originalUrl": c.OriginalURL(),
			"protocol":    c.Protocol(),
		}
	}

	if config.EnableRequest {
		response["request"] = map[string]interface{}{
			"query": c.Queries(),
			"body":  string(c.Body()),
			"path":  c.Path(),
			"ip":    c.IP(),
			"ips":   c.IPs(),
			"route": c.Route().Path,
		}
	}

	if config.EnableHeader {
		headers := make(map[string]string)
		c.Request().Header.VisitAll(func(key, value []byte) {
			headers[string(key)] = string(value)
		})
		response["headers"] = headers
	}

	if config.EnableCookies {
		cookies := make(map[string]string)
		c.Request().Header.VisitAllCookie(func(key, value []byte) {
			cookies[string(key)] = string(value)
		})
		response["cookies"] = cookies
	}

	if config.EnableEnvironment {
		env := make(map[string]string)
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			if len(pair) == 2 {
				env[pair[0]] = pair[1]
			}
		}
		response["environment"] = env
	}

	return response
}
