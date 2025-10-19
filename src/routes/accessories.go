package routes

import (
	"encoding/json"
	"fmt"
	"os"
	"skycrypt/src/api"
	"skycrypt/src/stats"
	"strings"
	"time"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	skyhelpernetworthgo "github.com/SkyCryptWebsite/SkyHelper-Networth-Go"
	"github.com/gofiber/fiber/v2"
)

// AccessoriesHandler godoc
// @Summary Get accessories stats of a specified player
// @Description Returns accessories for the given user and profile ID
// @Tags accessories
// @Accept  json
// @Produce  json
// @Param uuid path string true "User UUID"
// @Param profileId path string true "Profile ID"
// @Success 200 {object} models.GetMissingAccessoresOutput
// @Failure 400 {object} models.ProcessingError
// @Failure 500 {object} models.ProcessingError
// @Router /api/accessories/{uuid}/{profileId} [get]
func AccessoriesHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")

	profile, err := api.GetProfile(uuid, profileId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	userProfile := profile.Members[uuid]
	if userProfile.Inventory == nil {
		userProfile.Inventory = &skycrypttypes.Inventory{}
	}

	specifiedInventories := skyhelpernetworthgo.SpecifiedInventory{
		"talisman_bag": userProfile.Inventory.BagContents.TalismanBag,
	}

	decodedItems, err := skyhelpernetworthgo.CalculateFromSpecifiedInventories(specifiedInventories, skyhelpernetworthgo.NetworthOptions{
		IncludeItemData:  true,
		KeepInvalidItems: true,
	}.ToInternal())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to calculate items: %v", err),
		})
	}

	accessories := []*skycrypttypes.Item{}
	if decodedItems.Types["talisman_bag"] != nil {
		for _, item := range decodedItems.Types["talisman_bag"].Items {
			if item.ItemData != nil {
				item.ItemData.Price = item.Price
			}

			accessories = append(accessories, item.ItemData)
		}
	}

	items := map[string][]*skycrypttypes.Item{
		"talisman_bag": accessories,
	}

	disabledPacks := []string{""}
	disabledPacksCookies := c.Cookies("disabledPacks", "FAILED")
	if disabledPacksCookies != "FAILED" {
		var parsedPacks []string
		err := json.Unmarshal([]byte(disabledPacksCookies), &parsedPacks)
		if err == nil {
			disabledPacks = append(disabledPacks, parsedPacks...)
		}
	} else if os.Getenv("DEV") == "true" {
		disabledResourcePacks := c.Query("disabledPacks", "")
		if disabledResourcePacks != "" {
			disabledPacks = strings.Split(disabledResourcePacks, ",")
		}
	}

	fmt.Printf("Returning /api/accessories/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"accessories": stats.GetAccessories(&userProfile, items, disabledPacks),
	})
}
