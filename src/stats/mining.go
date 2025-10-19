package stats

import (
	"fmt"
	"os"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	statsItems "skycrypt/src/stats/items"
	stats "skycrypt/src/stats/leveling"

	"skycrypt/src/utility"
	"slices"
	"strings"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

func getPeakOfTheMountain(userProfile *skycrypttypes.Member) models.PeakOfTheMountain {
	return models.PeakOfTheMountain{
		Level:    userProfile.Mining.Nodes["special_0"],
		MaxLevel: constants.MAX_PEAK_OF_THE_MOUNTAIN_LEVEL,
	}
}

func getSelectedPickaxeAbility(userProfile *skycrypttypes.Member) string {
	if userProfile.Mining.SelectedPickaxeAbility == "" {
		return "None"
	}

	return userProfile.Mining.SelectedPickaxeAbility
}

func calcHotmTokens(hotmTier int, potmTier int) int {
	tokens := 0

	for tier := 1; tier <= hotmTier; tier++ {
		if reward, ok := constants.HOTM_REWARDS[tier]; ok {
			tokens += reward.TokenOfTheMountain
		}
	}

	for tier := 1; tier <= potmTier; tier++ {
		if reward, ok := constants.POTM_REWARDS[tier]; ok {
			tokens += reward.TokenOfTheMountain
		}
	}

	return tokens
}

func getHotmTokens(hotmLevel models.Skill, userProfile *skycrypttypes.Member) models.HotmTokens {
	potmTier := userProfile.Mining.Nodes["special_0"]
	hotmTokensAmount := calcHotmTokens(hotmLevel.Level, potmTier)
	return models.HotmTokens{
		Total:     hotmTokensAmount,
		Spent:     userProfile.Mining.TokensSpent,
		Available: hotmTokensAmount - userProfile.Mining.TokensSpent,
	}
}

func getCommissions(userProfile *skycrypttypes.Member, player *skycrypttypes.Player) models.Commissions {
	var milestone = 0
	if userProfile.Objectives == nil {
		return models.Commissions{
			Milestone:   milestone,
			Completions: player.Achievements.HotMCommissions,
		}
	}

	for _, tutorial := range userProfile.Objectives.Tutorial {
		if strings.HasPrefix(tutorial, "commission_milestone_reward_mining_xp_tier_") {
			tier := strings.Split(tutorial, "_")
			if len(tier) > 0 {
				lastElement := tier[len(tier)-1]
				parsedTier, err := utility.ParseInt(lastElement)
				if err == nil {
					if parsedTier > milestone {
						milestone = parsedTier
					}
				}
			}

		}
	}

	return models.Commissions{
		Milestone:   milestone,
		Completions: player.Achievements.HotMCommissions,
	}
}

func getCrystalHollows(userProfile *skycrypttypes.Member) models.CrystalHollows {
	totalRuns := 0
	for _, crystalData := range userProfile.Mining.Crystals {
		if crystalData.TotalPlaced > totalRuns {
			totalRuns = crystalData.TotalPlaced
		}
	}

	crystalHollows := models.CrystalHollows{
		CrystalHollowsLastAccess: userProfile.Mining.GreaterMinesLastAccess,
		NucleusRuns:              totalRuns,
		Progress: models.CrystalNucleusRuns{
			Crystals: make(map[string]string),
			Parts:    make(map[string]string),
		},
	}

	for _, crystal := range constants.GEMSTONE_CRYSTALS {
		crystalKey := crystal + "_crystal"

		if crystalData, exists := userProfile.Mining.Crystals[crystalKey]; exists {
			crystalHollows.Progress.Crystals[crystal] = crystalData.State
		} else {
			crystalHollows.Progress.Crystals[crystal] = "NOT_FOUND"
		}
	}

	for _, part := range constants.PRECURSOR_PARTS {
		partKey := strings.ToUpper(part)
		partsDelivered := userProfile.Mining.Biomes.Precursor.PartsDelivered

		if slices.Contains(partsDelivered, partKey) {
			crystalHollows.Progress.Parts[part] = "DELIVERED"
		} else {
			crystalHollows.Progress.Parts[part] = "NOT_DELIVERED"
		}
	}

	return crystalHollows
}

func getPowderAmount(userProfile *skycrypttypes.Member, powderType string) models.PowderAmount {
	spent := 0
	total := 0
	available := 0

	switch powderType {
	case "mithril":
		available = userProfile.Mining.PowderMithril
		spent = userProfile.Mining.PowderSpentMithril
		total = userProfile.Mining.PowderMithrilTotal
	case "gemstone":
		available = userProfile.Mining.PowderGemstone
		spent = userProfile.Mining.PowderSpentGemstone
		total = userProfile.Mining.PowderGemstoneTotal
	case "glacite":
		available = userProfile.Mining.PowderGlacite
		spent = userProfile.Mining.PowderSpentGlacite
		total = userProfile.Mining.PowderGlaciteTotal
	}

	return models.PowderAmount{
		Spent:     spent,
		Total:     total,
		Available: available,
	}
}

func getPowder(userProfile *skycrypttypes.Member) models.PowderOutput {
	return models.PowderOutput{
		Mithril:  getPowderAmount(userProfile, "mithril"),
		Gemstone: getPowderAmount(userProfile, "gemstone"),
		Glacite:  getPowderAmount(userProfile, "glacite"),
	}

}

func getForge(userProfile *skycrypttypes.Member) []models.ForgeOutput {
	output := []models.ForgeOutput{}

	quickForgeLevel := userProfile.Mining.Nodes["forge_time"]
	var quickForge float64
	if quickForgeLevel > 0 {
		if quickForgeLevel <= 19 {
			quickForge = float64(100+quickForgeLevel*5) / 10.0
		} else {
			quickForge = 300.0 / 10.0
		}
	} else {
		quickForge = 0
	}

	for _, item := range userProfile.Forge.ForgeProcesses.Forge {
		forgeConst := constants.FORGE[item.Id]
		duration := float64(forgeConst.Duration) - float64(forgeConst.Duration)*quickForge
		endingTime := item.StartTime + int64(float64(forgeConst.Duration)-float64(forgeConst.Duration)*(quickForge/100))
		output = append(output, models.ForgeOutput{
			ID:           item.Id,
			Name:         forgeConst.Name,
			Slot:         item.Slot,
			StartingTime: item.StartTime,
			EndingTime:   endingTime,
			Duration:     duration,
		})
	}

	return output
}

func getGlaciteTunnels(userProfile *skycrypttypes.Member) models.GlaciteTunnels {
	if userProfile.GlaciteTunnels == nil {
		return models.GlaciteTunnels{}
	}

	output := models.GlaciteTunnels{
		MineshaftsEntered: userProfile.GlaciteTunnels.MineshaftsEntered,
		FossilDust:        userProfile.GlaciteTunnels.FossilDust,
		Corpses:           models.Corpses{},
		Fossils:           models.Fossils{},
	}

	found := 0
	for corpseId, corpseTexture := range constants.CORPSES {
		found += userProfile.GlaciteTunnels.CorpsesLooted[corpseId]
		corpseData := models.Corpse{
			Amount:  userProfile.GlaciteTunnels.CorpsesLooted[corpseId],
			Name:    utility.TitleCase(corpseId),
			Texture: corpseTexture,
		}

		if os.Getenv("DEV") == "true" {
			corpseData.Texture = strings.Replace(corpseData.Texture, "/api/item/", "http://localhost:8080/api/item/", 1)
		}

		output.Corpses.Corpses = append(output.Corpses.Corpses, corpseData)
	}

	output.Corpses.Found = found
	output.Corpses.Max = len(constants.CORPSES)

	found = 0
	for _, fossil := range constants.FOSSILS {
		isFound := slices.Contains(userProfile.GlaciteTunnels.FossilsDonated, fossil)
		if isFound {
			found++
		}

		texture := fmt.Sprintf("/api/item/%s_FOSSIL", fossil)
		if fossil == "HELIX" {
			texture = fmt.Sprintf("/api/item/%s", fossil)
		}

		if os.Getenv("DEV") == "true" {
			texture = strings.Replace(texture, "/api/item/", "http://localhost:8080/api/item/", 1)
		}

		fossilData := models.Fossil{
			Name:    utility.TitleCase(fossil),
			Texture: texture,
			Found:   isFound,
		}

		output.Fossils.Fossils = append(output.Fossils.Fossils, fossilData)
	}

	output.Fossils.Found = found
	output.Fossils.Max = len(constants.FOSSILS)

	return output
}

func GetMining(userProfile *skycrypttypes.Member, player *skycrypttypes.Player, items []models.ProcessedItem) models.MiningOutput {
	HOTMLevel := stats.GetLevelByXp(int(userProfile.Mining.Experience), &stats.ExtraSkillData{Type: "hotm"})

	return models.MiningOutput{
		Level:                  HOTMLevel,
		PeakOfTheMountain:      getPeakOfTheMountain(userProfile),
		SelectedPickaxeAbility: getSelectedPickaxeAbility(userProfile),
		Tokens:                 getHotmTokens(HOTMLevel, userProfile),
		Commissions:            getCommissions(userProfile, player),
		CrystalHollows:         getCrystalHollows(userProfile),
		Powder:                 getPowder(userProfile),
		GlaciteTunnels:         getGlaciteTunnels(userProfile),
		Forge:                  getForge(userProfile),
		Tools:                  statsItems.GetSkillTools("mining", items),
		Hotm:                   []models.ProcessedItem{},
	}
}
