package routes

import (
	"fmt"
	"skycrypt/src/api"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	stats "skycrypt/src/stats"
	statsItems "skycrypt/src/stats/items"

	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

// GearHandler godoc
// @Summary Get gear stats of a specified player
// @Description Returns gear for the given user and profile ID
// @Tags gear
// @Accept  json
// @Produce  json
// @Param uuid path string true "User UUID"
// @Param profileId path string true "Profile ID"
// @Success 200 {object} models.Gear
// @Failure 400 {object} models.ProcessingError
// @Failure 500 {object} models.ProcessingError
// @Router /api/gear/{uuid}/{profileId} [get]
func GearHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")

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
		profile, err := api.GetProfile(uuid, profileId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to get profile: %v", err),
			})
		}

		userProfileValue := profile.Members[uuid]
		userProfile := &userProfileValue

		items, err = stats.GetItems(userProfile, profile.ProfileID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to get items: %v", err),
			})
		}
	}

	var processedItems = make(map[string][]models.ProcessedItem)
	inventoryKeys := []string{"armor", "equipment", "wardrobe", "inventory", "enderchest", "backpack"}
	for _, inventoryId := range inventoryKeys {
		inventoryData := items[inventoryId]
		if len(inventoryData) == 0 {
			continue
		}

		processedItems[inventoryId] = statsItems.ProcessItems(&inventoryData, inventoryId)
	}

	allItems := make([]models.ProcessedItem, 0)
	for _, inventoryId := range inventoryKeys {
		if processedItems[inventoryId] != nil {
			allItems = append(allItems, processedItems[inventoryId]...)
		}
	}

	fmt.Printf("Returning /api/gear/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"gear": stats.GetGear(processedItems, allItems),
	})
}
