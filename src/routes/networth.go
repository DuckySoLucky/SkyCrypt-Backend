package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/stats"
	"time"

	skyhelpernetworthgo "github.com/SkyCryptWebsite/SkyHelper-Networth-Go"
	"github.com/gofiber/fiber/v2"
)

type NetworthResult struct {
	Networth            float64                  `json:"networth"`
	UnsoulboundNetworth float64                  `json:"unsoulboundNetworth"`
	NoInventory         bool                     `json:"noInventory"`
	IsNonCosmetic       bool                     `json:"isNonCosmetic"`
	Purse               float64                  `json:"purse"`
	Bank                float64                  `json:"bank"`
	PersonalBank        float64                  `json:"personalBank"`
	Types               map[string]*NetworthType `json:"types"`
}

type NetworthType struct {
	Total            float64 `json:"total"`
	UnsoulboundTotal float64 `json:"unsoulboundTotal"`
}

// NetworthHandler godoc
// @Summary Get networth of a specified player
// @Description Returns networth for the given user and profile ID
// @Tags networth
// @Produce  json
// @Param uuid path string true "User UUID"
// @Param profileId path string true "Profile ID"
// @Success 200 {object} map[string]NetworthResult
// @Failure 400 {object} models.ProcessingError
// @Failure 500 {object} models.ProcessingError
// @Router /api/networth/{uuid}/{profileId} [get]
func NetworthHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")
	if len(profileId) > 0 && profileId[0] == '/' {
		profileId = profileId[1:]
	}

	mowojang, err := api.ResolvePlayer(uuid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to resolve player: %v", err),
		})
	}

	profiles, err := api.GetProfiles(mowojang.UUID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	profile, err := stats.GetProfile(profiles, profileId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	profileMuseum, err := api.GetMuseum(profile.ProfileID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get museum: %v", err),
		})
	}
	userProfileValue := profile.Members[mowojang.UUID]
	museum := profileMuseum[mowojang.UUID]
	userProfile := &userProfileValue

	var bankBalance float64
	if profile.Banking != nil && profile.Banking.Balance != nil {
		bankBalance = *profile.Banking.Balance
	} else {
		bankBalance = 0.0
	}

	calculator, err := skyhelpernetworthgo.NewProfileNetworthCalculator(userProfile, museum, bankBalance)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create networth calculator: %v", err),
		})
	}

	networth := calculator.GetNetworth()
	nonCosmeticNetworth := calculator.GetNonCosmeticNetworth()
	formattedNetworth := map[string]float64{
		"normal":      networth.Networth,
		"nonCosmetic": nonCosmeticNetworth.Networth,
	}

	go stats.StoreEmbedData(mowojang, userProfile, profile, formattedNetworth)

	fmt.Printf("Returning /api/networth/%s in %s\n", uuid, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"normal":      networth,
		"nonCosmetic": nonCosmeticNetworth,
	})

}
