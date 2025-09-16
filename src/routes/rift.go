package routes

import (
	"fmt"
	"skycrypt/src/api"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	"skycrypt/src/stats"
	statsitems "skycrypt/src/stats/items"
	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func RiftHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")

	profile, err := api.GetProfile(uuid, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	userProfileValue := profile.Members[uuid]
	userProfile := &userProfileValue

	var items map[string][]skycrypttypes.Item
	cache, err := redis.Get(fmt.Sprintf("items:%s", profileId))
	if err == nil && cache != "" {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		err = json.Unmarshal([]byte(cache), &items)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to parse items: %v", err),
			})
		}
	} else {
		items, err = stats.GetItems(userProfile, profile.ProfileID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to get items: %v", err),
			})
		}
	}

	var processedItems = make(map[string][]models.ProcessedItem)
	inventoryKeys := []string{"rift_armor", "rift_equipment"}
	for _, inventoryId := range inventoryKeys {
		inventoryData := items[inventoryId]
		if len(inventoryData) == 0 {
			continue
		}

		processedItems[inventoryId] = statsitems.ProcessItems(&inventoryData, inventoryId)
	}

	fmt.Printf("Returning /api/rift/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"rift": stats.GetRift(userProfile, processedItems),
	})
}
