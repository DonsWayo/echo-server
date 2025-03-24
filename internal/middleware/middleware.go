package middleware

import (
	"strconv"
	"strings"
	"time"

	"github.com/a-safe-digital/echo-server/internal/config"
	"github.com/a-safe-digital/echo-server/internal/platform/utils"
	"github.com/gofiber/fiber/v3"
)

func ProcessRequest(config *config.Configuration) fiber.Handler {
	return func(c fiber.Ctx) error {
		statusCode := processStatusCode(c, config)
		processHeaders(c, config)
		processWaitTime(c, config)

		if statusCode > 0 {
			c.Status(statusCode)
		}

		return c.Next()
	}
}

func processStatusCode(c fiber.Ctx, config *config.Configuration) int {
	codeStr := utils.GetValueFromHeaderOrQuery(c, config.Commands.HTTPCode.Header, config.Commands.HTTPCode.Query)
	if codeStr == "" {
		return 0
	}
	if strings.Contains(codeStr, "-") {
		codes := strings.Split(codeStr, "-")
		if len(codes) > 0 {
			randomIndex := time.Now().UnixNano() % int64(len(codes))
			codeStr = codes[randomIndex]
		}
	}
	code, err := strconv.Atoi(codeStr)
	if err != nil || code < 200 || code > 599 {
		return 0
	}
	return code
}

func processHeaders(c fiber.Ctx, config *config.Configuration) {
	if !config.EnableHeader {
		return
	}
	headersStr := utils.GetValueFromHeaderOrQuery(c, config.Commands.HTTPHeaders.Header, config.Commands.HTTPHeaders.Query)
	if headersStr == "" {
		return
	}
	headerPairs := strings.Split(headersStr, ",")
	for _, pair := range headerPairs {
		pair = strings.TrimSpace(pair)
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			c.Set(parts[0], parts[1])
		}
	}
}

func processWaitTime(c fiber.Ctx, config *config.Configuration) {
	timeStr := utils.GetValueFromHeaderOrQuery(c, config.Commands.Time.Header, config.Commands.Time.Query)
	if timeStr == "" {
		return
	}
	waitTime, err := strconv.Atoi(timeStr)
	if err != nil {
		return
	}
	if waitTime < config.Controls.Times.Min {
		waitTime = config.Controls.Times.Min
	} else if waitTime > config.Controls.Times.Max {
		waitTime = config.Controls.Times.Max
	}
	if waitTime > 0 {
		time.Sleep(time.Duration(waitTime) * time.Millisecond)
	}
}
