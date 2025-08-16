package routes

import (
	"github.com/gofiber/fiber/v2"
)

func NetworthHandler(c *fiber.Ctx) error {
	/*timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")
	if len(profileId) > 0 && profileId[0] == '/' {
		profileId = profileId[1:]
	}

	mowojang, err := api.ResolvePlayer(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to resolve player: %v", err),
		})
	}

	profiles, err := api.GetProfiles(mowojang.UUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	profile, err := stats.GetProfile(profiles, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	profileMuseum, err := api.GetMuseum(profile.ProfileID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get museum: %v", err),
		})
	}
	userProfileValue := profile.Members[mowojang.UUID]
	museum := profileMuseum[mowojang.UUID]
	userProfile := &userProfileValue

	calculator, err := skyhelpernetworthgo.NewProfileNetworthCalculator(userProfile, museum, *profile.Banking.Balance)
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
		"networth": map[string]any{
			"normal":      networth,
			"nonCosmetic": nonCosmeticNetworth,
		},
	})*/
	return c.JSON(fiber.Map{
		"networth": nil,
	})
}
