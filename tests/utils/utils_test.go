package utils_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/a-safe-digital/echo-server/internal/platform/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetValueFromHeaderOrQuery(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c fiber.Ctx) error {
		value := utils.GetValueFromHeaderOrQuery(c, "X-Test-Header", "test_query")
		return c.SendString(value)
	})

	tests := []struct {
		name       string
		headerKey  string
		headerVal  string
		queryKey   string
		queryVal   string
		expectVal  string
	}{
		{
			name:      "Get value from header",
			headerKey: "X-Test-Header",
			headerVal: "header-value",
			queryKey:  "test_query",
			queryVal:  "",
			expectVal: "header-value",
		},
		{
			name:      "Get value from query",
			headerKey: "X-Test-Header",
			headerVal: "",
			queryKey:  "test_query",
			queryVal:  "query-value",
			expectVal: "query-value",
		},
		{
			name:      "Header takes precedence over query",
			headerKey: "X-Test-Header",
			headerVal: "header-value",
			queryKey:  "test_query",
			queryVal:  "query-value",
			expectVal: "header-value",
		},
		{
			name:      "No value in header or query",
			headerKey: "X-Test-Header",
			headerVal: "",
			queryKey:  "test_query",
			queryVal:  "",
			expectVal: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/test"
			if tt.queryVal != "" {
				url = url + "?" + tt.queryKey + "=" + tt.queryVal
			}
			
			req := httptest.NewRequest("GET", url, nil)
			if tt.headerVal != "" {
				req.Header.Set(tt.headerKey, tt.headerVal)
			}
			
			resp, err := app.Test(req)
			assert.NoError(t, err)
			
			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			
			assert.Equal(t, tt.expectVal, string(body))
		})
	}
}
