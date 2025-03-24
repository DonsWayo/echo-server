package config_test

import (
	"os"
	"testing"

	"github.com/a-safe-digital/echo-server/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultConfig(t *testing.T) {
	cfg := config.NewDefaultConfig()
	
	assert.Equal(t, "80", cfg.Port)
	assert.True(t, cfg.EnableHost)
	assert.True(t, cfg.EnableHTTP)
	assert.True(t, cfg.EnableFile)
	assert.False(t, cfg.LogsIgnorePing)
}

func TestEnvironmentVariableOverrides(t *testing.T) {
	originalPort := os.Getenv("PORT")
	originalEnableHost := os.Getenv("ENABLE__HOST")
	
	defer func() {
		os.Setenv("PORT", originalPort)
		os.Setenv("ENABLE__HOST", originalEnableHost)
	}()
	
	os.Setenv("PORT", "8080")
	os.Setenv("ENABLE__HOST", "false")
	
	cfg := config.NewDefaultConfig()
	
	assert.Equal(t, "8080", cfg.Port)
	assert.False(t, cfg.EnableHost)
}

func TestCommandsConfig(t *testing.T) {
	cmds := config.NewDefaultCommandsConfig()
	
	assert.Equal(t, "X-ECHO-CODE", cmds.HTTPCode.Header)
	assert.Equal(t, "echo_code", cmds.HTTPCode.Query)
	
	assert.Equal(t, "X-ECHO-BODY", cmds.HTTPBody.Header)
	assert.Equal(t, "echo_body", cmds.HTTPBody.Query)
	
	assert.Equal(t, "X-ECHO-HEADER", cmds.HTTPHeaders.Header)
	assert.Equal(t, "echo_header", cmds.HTTPHeaders.Query)
	
	assert.Equal(t, "X-ECHO-FILE", cmds.File.Header)
	assert.Equal(t, "echo_file", cmds.File.Query)
	
	assert.Equal(t, "X-ECHO-ENV", cmds.HTTPEnvBody.Header)
	assert.Equal(t, "echo_env", cmds.HTTPEnvBody.Query)
	
	assert.Equal(t, "X-ECHO-TIME", cmds.Time.Header)
	assert.Equal(t, "echo_time", cmds.Time.Query)
}

func TestControlsConfig(t *testing.T) {
	controls := config.NewDefaultControlsConfig()
	
	assert.Equal(t, 0, controls.Times.Min)
	assert.Equal(t, 60000, controls.Times.Max)
}
