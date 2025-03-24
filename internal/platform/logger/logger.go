package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/a-safe-digital/echo-server/internal/config"
	"github.com/a-safe-digital/echo-server/internal/platform/constants"
	"github.com/gofiber/fiber/v3"
)

func Logger(config *config.Configuration) fiber.Handler {
	return func(c fiber.Ctx) error {
		if config.LogsIgnorePing && c.Path() == constants.RouteHealth {
			return c.Next()
		}

		start := time.Now()
		
		err := c.Next()
		
		latency := time.Since(start)
		log.Printf(
			"[%s] %s - %d - %s - %s",
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			latency,
			c.IP(),
		)
		
		return err
	}
}

func PrintServerInfo(config *config.Configuration) {
	fmt.Println(`
    ______      __             _____                          
   / ____/___  / /_  ____     / ___/___  ______   _____  _____
  / __/ / __ \/ __ \/ __ \    \__ \/ _ \/ ___/ | / / _ \/ ___/
 / /___/ /_/ / / / / /_/ /   ___/ /  __/ /   | |/ /  __/ /    
/_____/\____/_/ /_/\____/   /____/\___/_/    |___/\___/_/     
--------------------------------------------------`)
	fmt.Printf("INFO Server started on: \thttp://127.0.0.1:%s\n", config.Port)
	fmt.Printf("INFO Prefork: \t\tDisabled\n")
	fmt.Printf("INFO PID: \t\t%d\n", os.Getpid())
}
