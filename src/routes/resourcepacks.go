package routes

import (
	"encoding/json"
	"fmt"
	"os"
	"skycrypt/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

var RESOURCE_PACKS = []models.ResourcePackConfig{}

// ResourcePackHandler godoc
// @Summary Get list of resource packs
// @Description Returns a list of resource packs
// @Tags resourcepacks
// @Accept  json
// @Produce  json
// @Success 200 {object} []ResourcePackConfig
// @Router /api/resourcepacks/{uuid}/{profileId} [get]
func ResourcePackHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	if len(RESOURCE_PACKS) == 0 {
		filePath := "assets/resourcepacks/"
		files, err := os.ReadDir(filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read resource packs directory",
			})
		}

		for _, file := range files {
			if !file.IsDir() {
				continue
			}

			configPath := fmt.Sprintf("%s/%s/config.json", filePath, file.Name())
			configFile, err := os.Open(configPath)
			if err != nil {
				continue
			}

			defer configFile.Close()

			var configData models.ResourcePackConfig
			if err := json.NewDecoder(configFile).Decode(&configData); err != nil {
				continue
			}

			configData.Icon = fmt.Sprintf("/assets/resourcepacks/%s/pack.png", file.Name())
			RESOURCE_PACKS = append(RESOURCE_PACKS, configData)
		}
	}

	fmt.Printf("Returning /api/resourcepacks in %s\n", time.Since(timeNow))

	return c.JSON(fiber.Map{
		"resourcepacks": RESOURCE_PACKS,
	})
}
