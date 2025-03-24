package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Configuration struct {
	Port              string
	EnableHost        bool
	EnableHTTP        bool
	EnableRequest     bool
	EnableCookies     bool
	EnableHeader      bool
	EnableEnvironment bool
	EnableFile        bool
	LogsIgnorePing    bool
	Controls          struct {
		Times struct {
			Min int
			Max int
		}
	}
	Commands struct {
		HTTPBody struct {
			Query  string
			Header string
		}
		HTTPEnvBody struct {
			Query  string
			Header string
		}
		HTTPCode struct {
			Query  string
			Header string
		}
		HTTPHeaders struct {
			Query  string
			Header string
		}
		Time struct {
			Query  string
			Header string
		}
		File struct {
			Query  string
			Header string
		}
	}
}

func NewDefaultConfig() *Configuration {
	config := &Configuration{
		Port:              getEnv("PORT", "80"),
		EnableHost:        getEnvBool("ENABLE__HOST", true),
		EnableHTTP:        getEnvBool("ENABLE__HTTP", true),
		EnableRequest:     getEnvBool("ENABLE__REQUEST", true),
		EnableCookies:     getEnvBool("ENABLE__COOKIES", true),
		EnableHeader:      getEnvBool("ENABLE__HEADER", true),
		EnableEnvironment: getEnvBool("ENABLE__ENVIRONMENT", true),
		EnableFile:        getEnvBool("ENABLE__FILE", true),
		LogsIgnorePing:    getEnvBool("LOGS__IGNORE__PING", false),
	}

	config.Controls.Times.Min = getEnvInt("CONTROLS__TIMES__MIN", 0)
	config.Controls.Times.Max = getEnvInt("CONTROLS__TIMES__MAX", 60000)

	config.Commands.HTTPBody.Query = getEnv("COMMANDS__HTTPBODY__QUERY", "echo_body")
	config.Commands.HTTPBody.Header = getEnv("COMMANDS__HTTPBODY__HEADER", "x-echo-body")

	config.Commands.HTTPEnvBody.Query = getEnv("COMMANDS__HTTPENVBODY__QUERY", "echo_env_body")
	config.Commands.HTTPEnvBody.Header = getEnv("COMMANDS__HTTPENVBODY__HEADER", "x-echo-env-body")

	config.Commands.HTTPCode.Query = getEnv("COMMANDS__HTTPCODE__QUERY", "echo_code")
	config.Commands.HTTPCode.Header = getEnv("COMMANDS__HTTPCODE__HEADER", "x-echo-code")

	config.Commands.HTTPHeaders.Query = getEnv("COMMANDS__HTTPHEADERS__QUERY", "echo_header")
	config.Commands.HTTPHeaders.Header = getEnv("COMMANDS__HTTPHEADERS__HEADER", "x-echo-header")

	config.Commands.Time.Query = getEnv("COMMANDS__TIME__QUERY", "echo_time")
	config.Commands.Time.Header = getEnv("COMMANDS__TIME__HEADER", "x-echo-time")

	config.Commands.File.Query = getEnv("COMMANDS__FILE__QUERY", "echo_file")
	config.Commands.File.Header = getEnv("COMMANDS__FILE__HEADER", "x-echo-file")

	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func main() {
	config := NewDefaultConfig()
	app := fiber.New()
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

	app.Use(func(c fiber.Ctx) error {
		statusCode := processStatusCode(c, config)
		processHeaders(c, config)
		processWaitTime(c, config)
		if statusCode > 0 {
			c.Status(statusCode)
		}

		return c.Next()
	})

	app.Get("/", func(c fiber.Ctx) error {
		if filePath := getValueFromHeaderOrQuery(c, config.Commands.File.Header, config.Commands.File.Query); filePath != "" && config.EnableFile {
			return handleFileExploration(c, filePath)
		}
		if envVar := getValueFromHeaderOrQuery(c, config.Commands.HTTPEnvBody.Header, config.Commands.HTTPEnvBody.Query); envVar != "" && config.EnableEnvironment {
			return c.JSON(os.Getenv(envVar))
		}
		if body := getValueFromHeaderOrQuery(c, config.Commands.HTTPBody.Header, config.Commands.HTTPBody.Query); body != "" {
			return c.JSON(body)
		}
		response := buildDefaultResponse(c, config)
		return c.JSON(response)
	})

	app.All("/", func(c fiber.Ctx) error {
		body := c.Body()
		if len(body) > 0 {
			return c.Send(body)
		}
		if filePath := getValueFromHeaderOrQuery(c, config.Commands.File.Header, config.Commands.File.Query); filePath != "" && config.EnableFile {
			return handleFileExploration(c, filePath)
		}
		if envVar := getValueFromHeaderOrQuery(c, config.Commands.HTTPEnvBody.Header, config.Commands.HTTPEnvBody.Query); envVar != "" && config.EnableEnvironment {
			return c.JSON(os.Getenv(envVar))
		}
		if body := getValueFromHeaderOrQuery(c, config.Commands.HTTPBody.Header, config.Commands.HTTPBody.Query); body != "" {
			return c.JSON(body)
		}
		response := buildDefaultResponse(c, config)
		return c.JSON(response)
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	port := fmt.Sprintf(":%s", config.Port)
	log.Fatal(app.Listen(port))
}

func getValueFromHeaderOrQuery(c fiber.Ctx, headerKey, queryKey string) string {

	value := c.Get(headerKey)
	if value != "" {
		return value
	}

	return c.Query(queryKey)
}

func processStatusCode(c fiber.Ctx, config *Configuration) int {
	codeStr := getValueFromHeaderOrQuery(c, config.Commands.HTTPCode.Header, config.Commands.HTTPCode.Query)
	if codeStr == "" {
		return 0
	}
	if strings.Contains(codeStr, "-") {
		codes := strings.Split(codeStr, "-")
		if len(codes) > 0 {
			randomIndex := rand.Intn(len(codes))
			codeStr = codes[randomIndex]
		}
	}
	code, err := strconv.Atoi(codeStr)
	if err != nil || code < 200 || code > 599 {
		return 0
	}
	return code
}

func processHeaders(c fiber.Ctx, config *Configuration) {
	if !config.EnableHeader {
		return
	}
	headersStr := getValueFromHeaderOrQuery(c, config.Commands.HTTPHeaders.Header, config.Commands.HTTPHeaders.Query)
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

func processWaitTime(c fiber.Ctx, config *Configuration) {
	timeStr := getValueFromHeaderOrQuery(c, config.Commands.Time.Header, config.Commands.Time.Query)
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

func handleFileExploration(c fiber.Ctx, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return c.Status(404).JSON(map[string]string{"error": "Path not found"})
	}

	if info.IsDir() {
		var files []string
		err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if filepath.Dir(path) == filepath.Clean(path) && path != filepath.Clean(path) {
				files = append(files, d.Name())
			}
			return fs.SkipDir
		})
		if err != nil {
			return c.Status(500).JSON(map[string]string{"error": err.Error()})
		}

		return c.JSON(files)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return c.Status(500).JSON(map[string]string{"error": err.Error()})
	}

	var jsonObj interface{}
	if json.Unmarshal(content, &jsonObj) == nil {
		return c.JSON(jsonObj)
	}

	return c.SendString(string(content))
}

func buildDefaultResponse(c fiber.Ctx, config *Configuration) map[string]interface{} {
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
