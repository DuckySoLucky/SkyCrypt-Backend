package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/stats"
	"time"

	"github.com/gofiber/fiber/v2"
)

// PlayerStatsHandler godoc
// @Summary Get player stats of a specified player
// @Description Returns player stats for the given user and profile ID
// @Tags playerStats
// @Accept  json
// @Produce  json
// @Param uuid path string true "User UUID"
// @Param profileId path string true "Profile ID"
// @Success 200 {object} map[string]models.StatsInfo
// @Failure 400 {object} models.ProcessingError
// @Router /api/playerStats/{uuid}/{profileId} [get]
func PlayerStatsHandler(c *fiber.Ctx) error {
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

	fmt.Printf("Returning /api/playerStats/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"stats": stats.GetPlayerStats(userProfile, profile, profileId),
	})
}
