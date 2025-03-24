package handlers

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/a-safe-digital/echo-server/internal/platform/constants"
	"github.com/gofiber/fiber/v3"
)

func handleFileExploration(c fiber.Ctx, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return c.Status(constants.StatusNotFound).JSON(map[string]string{"error": constants.MessageNotFound})
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
			return c.Status(constants.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
		}

		return c.JSON(files)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return c.Status(constants.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	var jsonObj interface{}
	if json.Unmarshal(content, &jsonObj) == nil {
		return c.JSON(jsonObj)
	}

	return c.SendString(string(content))
}
