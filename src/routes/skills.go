package routes

import (
	"fmt"
	"skycrypt/src/api"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	"skycrypt/src/stats"
	statsItems "skycrypt/src/stats/items"

	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func SkillsHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")

	profile, err := api.GetProfile(uuid, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	player, err := api.GetPlayer(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get player: %v", err),
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

	fmt.Printf("Returning /api/skills/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"mining":     stats.GetMining(userProfile, player, allItems),
		"farming":    stats.GetFarming(userProfile, allItems),
		"fishing":    stats.GetFishing(userProfile, allItems),
		"enchanting": stats.GetEnchanting(userProfile),
	})
}
