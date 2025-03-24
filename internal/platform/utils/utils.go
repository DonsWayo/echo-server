package utils

import (
	"github.com/gofiber/fiber/v3"
)

func GetValueFromHeaderOrQuery(c fiber.Ctx, headerKey, queryKey string) string {
	value := c.Get(headerKey)
	if value != "" {
		return value
	}
	return c.Query(queryKey)
}
