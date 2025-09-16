package routes

import (
	"fmt"
	"skycrypt/src/api"
	redis "skycrypt/src/db"
	"skycrypt/src/stats"
	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func AccessoriesHandler(c *fiber.Ctx) error {
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

	fmt.Printf("Returning /api/accessories/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"accessories": stats.GetAccessories(userProfile, items),
	})
}
