package version

import (
	"fmt"
	"runtime"
)

var (
	Version = "1.0.0"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

func Info() string {
	return fmt.Sprintf(
		"Version: %s\nBuild Time: %s\nGit Commit: %s\nGo Version: %s\nOS/Arch: %s/%s",
		Version,
		BuildTime,
		GitCommit,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}

func GetVersion() string {
	return Version
}

func GetBuildInfo() map[string]string {
	return map[string]string{
		"version":    Version,
		"buildTime":  BuildTime,
		"gitCommit":  GitCommit,
		"goVersion":  runtime.Version(),
		"goOS":       runtime.GOOS,
		"goArch":     runtime.GOARCH,
	}
}
