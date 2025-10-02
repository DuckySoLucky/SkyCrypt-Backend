package stats

import (
	"os"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	statsitems "skycrypt/src/stats/items"
	"slices"
	"strings"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

func getMotes(userProfile *skycrypttypes.Member) models.RiftMotesOutput {
	return models.RiftMotesOutput{
		Purse:    int(userProfile.Currencies.MotesPurse),
		Lifetime: int(userProfile.PlayerStats.Rift.LifetimeMotesCollected),
		Orbs:     int(userProfile.PlayerStats.Rift.MotesOrbPickup),
	}
}

func getEnigma(userProfile *skycrypttypes.Member) models.RiftEnigmaOutput {

	return models.RiftEnigmaOutput{
		Souls:      len(userProfile.Rift.Enigma.FoundSouls),
		TotalSouls: constants.RIFT_ENIGMA_SOULS,
	}
}

func getCastle(userProfile *skycrypttypes.Member) models.RiftCastleOutput {
	return models.RiftCastleOutput{
		GrubberStacks: userProfile.Rift.Castle.GrubberStacks,
		MaxBurgers:    constants.RIFT_MAX_GRUBBER_STACKS,
	}
}

func getPorhtals(userProfile *skycrypttypes.Member) models.RiftPortalsOutput {
	porhtals, found := make([]models.RiftPorhtal, 0, len(userProfile.Rift.WitherCage.KilledEyes)), 0
	for _, portal := range constants.RIFT_EYES {
		isFound := slices.Contains(userProfile.Rift.WitherCage.KilledEyes, portal.Id)
		if isFound {
			found++
		}

		porhtalData := models.RiftPorhtal{
			Name:     portal.Name,
			Texture:  portal.Texture,
			Unlocked: isFound,
		}

		if os.Getenv("DEV") == "true" {
			porhtalData.Texture = strings.Replace(porhtalData.Texture, "/api/head/", "http://localhost:8080/api/head/", 1)
		}

		porhtals = append(porhtals, porhtalData)

	}

	return models.RiftPortalsOutput{
		PorhtalsFound: found,
		Porhtals:      porhtals,
	}
}

func getTimecharms(userProfile *skycrypttypes.Member) models.RiftTimecharmsOutput {
	timecharms := make([]models.RiftTimecharms, 0, len(constants.RIFT_TIMECHARMS))
	found := 0

	for _, charm := range constants.RIFT_TIMECHARMS {
		isFound, timestamp := false, int64(0)
		for _, id := range userProfile.Rift.Gallery.SecuredTrophies {
			if id.Type == charm.ID {
				isFound = true
				timestamp = id.Timestamp
				found++
				break
			}
		}

		timecharmData := models.RiftTimecharms{
			Name:       charm.Name,
			Texture:    charm.Texture,
			Unlocked:   isFound,
			UnlockedAt: timestamp,
		}

		if os.Getenv("DEV") == "true" {
			timecharmData.Texture = strings.Replace(timecharmData.Texture, "/api/item/", "http://localhost:8080/api/item/", 1)
		}

		timecharms = append(timecharms, timecharmData)
	}

	return models.RiftTimecharmsOutput{
		TimecharmsFound: found,
		Timecharms:      timecharms,
	}
}

func GetRift(userProfile *skycrypttypes.Member, items map[string][]models.ProcessedItem) *models.RiftOutput {
	return &models.RiftOutput{
		Visits:     int(userProfile.PlayerStats.Rift.Visits),
		Motes:      getMotes(userProfile),
		Enigma:     getEnigma(userProfile),
		Castle:     getCastle(userProfile),
		Porhtals:   getPorhtals(userProfile),
		Timecharms: getTimecharms(userProfile),
		Armor:      statsitems.GetArmor(items["rift_armor"]),
		Equipment:  statsitems.GetEquipment(items["rift_equipment"]),
	}
}
