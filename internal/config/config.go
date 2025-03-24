package config

import (
	"os"
	"strconv"
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
	Controls          ControlsConfig
	Commands          CommandsConfig
}

func NewDefaultConfig() *Configuration {
	return &Configuration{
		Port:              getEnv("PORT", "80"),
		EnableHost:        getEnvBool("ENABLE__HOST", true),
		EnableHTTP:        getEnvBool("ENABLE__HTTP", true),
		EnableRequest:     getEnvBool("ENABLE__REQUEST", true),
		EnableCookies:     getEnvBool("ENABLE__COOKIES", true),
		EnableHeader:      getEnvBool("ENABLE__HEADER", true),
		EnableEnvironment: getEnvBool("ENABLE__ENVIRONMENT", true),
		EnableFile:        getEnvBool("ENABLE__FILE", true),
		LogsIgnorePing:    getEnvBool("LOGS__IGNORE__PING", false),
		Controls:          NewDefaultControlsConfig(),
		Commands:          NewDefaultCommandsConfig(),
	}
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
