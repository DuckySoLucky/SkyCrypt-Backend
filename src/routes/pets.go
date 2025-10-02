package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/stats"
	"time"

	"github.com/gofiber/fiber/v2"
)

// PetsHandler godoc
// @Summary Get pets stats of a specified player
// @Description Returns pets for the given user and profile ID
// @Tags pets
// @Accept  json
// @Produce  json
// @Param uuid path string true "User UUID"
// @Param profileId path string true "Profile ID"
// @Success 200 {object} models.OutputPets
// @Failure 400 {object} models.ProcessingError
// @Router /api/pets/{uuid}/{profileId} [get]
func PetsHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")

	profile, err := api.GetProfile(uuid, profileId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	userProfileValue := profile.Members[uuid]
	userProfile := &userProfileValue

	fmt.Printf("Returning /api/pets/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"pets": stats.GetPets(userProfile, profile),
	})
}
