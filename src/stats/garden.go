package stats

import (
	"fmt"
	notenoughupdates "skycrypt/src/NotEnoughUpdates"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	stats "skycrypt/src/stats/leveling"
	"skycrypt/src/utility"
	"slices"
	"strings"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

type Garden struct {
	Level          models.Skill    `json:"level"`
	Visitors       Visitors        `json:"visitors"`
	CropMilestones []CropMilestone `json:"cropMilestones"`
	CropUpgrades   []CropUpgrade   `json:"cropUpgrades"`
	Composter      map[string]int  `json:"composter"`
	Plot           PlotLayout      `json:"plot"`
}

type Visitors struct {
	Visited        int                          `json:"visited"`
	Completed      int                          `json:"completed"`
	UniqueVisitors int                          `json:"uniqueVisitors"`
	Visitors       map[string]VisitorRarityData `json:"visitors"`
}

type VisitorRarityData struct {
	Visited   int `json:"visited"`
	Completed int `json:"completed"`
	Unique    int `json:"unique"`
	MaxUnique int `json:"maxUnique"`
}

type CropMilestone struct {
	Name    string       `json:"name"`
	Texture string       `json:"texture"`
	Level   models.Skill `json:"level"`
}

type CropUpgrade struct {
	Name    string       `json:"name"`
	Texture string       `json:"texture"`
	Level   models.Skill `json:"level"`
}

type PlotLayout struct {
	Unlocked int                    `json:"unlocked"`
	Total    int                    `json:"total"`
	BarnSkin string                 `json:"barnSkin"`
	Layout   []models.ProcessedItem `json:"layout"`
}

func getVisitors(gardenData *skycrypttypes.Garden) Visitors {
	VISITOR_RARITIES := notenoughupdates.NEUConstants.Garden.Visitors
	MAX_VISITORS := notenoughupdates.NEUConstants.Garden.MaxVisitors

	visited, completed, unique := 0, 0, map[string]bool{}
	visitors := make(map[string]VisitorRarityData, len(gardenData.CommissionData.Visits))
	for visitorId, amount := range gardenData.CommissionData.Visits {
		completed += gardenData.CommissionData.Completed[visitorId]
		unique[visitorId] = true
		visited += amount

		visitorData := visitors[VISITOR_RARITIES[visitorId]]
		if visitorData.MaxUnique == 0 {
			visitorData = VisitorRarityData{
				MaxUnique: MAX_VISITORS[VISITOR_RARITIES[visitorId]],
			}
		}

		visitorData.Unique += 1
		visitorData.Visited += amount
		visitorData.Completed += gardenData.CommissionData.Completed[visitorId]

		visitors[VISITOR_RARITIES[visitorId]] = visitorData

	}

	return Visitors{
		Visited:        visited,
		Completed:      completed,
		UniqueVisitors: len(unique),
		Visitors:       visitors,
	}
}

func getCropMilestones(gardenData *skycrypttypes.Garden) []CropMilestone {
	milestones := make([]CropMilestone, 0, len(gardenData.ResourcesCollected))
	for cropId, cropName := range constants.CROPS {
		milestones = append(milestones, CropMilestone{
			Name:    cropName,
			Texture: fmt.Sprintf("http://localhost:8080/api/item/%s", cropId),
			Level: stats.GetLevelByXp(int(gardenData.ResourcesCollected[cropId]), &stats.ExtraSkillData{
				Type: fmt.Sprintf("crop_milestone_%s", constants.CROP_TO_ID[cropId]),
			}),
		})
	}

	return milestones
}

func getCropUpgrades(gardenData *skycrypttypes.Garden) []CropUpgrade {
	upgrades := make([]CropUpgrade, 0, len(gardenData.CropUpgradeLevels))
	for cropId, cropName := range constants.CROPS {
		experience := stats.GetSkillExperience("crop_upgrade", int(gardenData.CropUpgradeLevels[cropId]))

		upgrades = append(upgrades, CropUpgrade{
			Name:    cropName,
			Texture: fmt.Sprintf("http://localhost:8080/api/item/%s", cropId),
			Level: stats.GetLevelByXp(experience, &stats.ExtraSkillData{
				Type: "crop_upgrade",
			}),
		})
	}

	return upgrades
}

func getComposter(gardenData *skycrypttypes.Garden) map[string]int {
	output := make(map[string]int, len(gardenData.ComposterData.Upgrades))
	for _, upgrade := range notenoughupdates.NEUConstants.Garden.ComposterUpgrades {
		output[upgrade] = int(gardenData.ComposterData.Upgrades[upgrade])
	}

	return output
}

func getPlotLayout(gardenData *skycrypttypes.Garden) PlotLayout {
	PLOT_LAYOUT := notenoughupdates.NEUConstants.Garden.SortedPlots
	PLOT_NAMES := notenoughupdates.NEUConstants.Garden.Plots

	output := PlotLayout{
		Unlocked: len(gardenData.UnlockedPlotsIds),
		Total:    len(PLOT_LAYOUT),
		BarnSkin: "",
		Layout:   make([]models.ProcessedItem, 0, len(PLOT_LAYOUT)),
	}

	for index, plot := range PLOT_LAYOUT {
		checkPlots := []string{}

		if index-5 >= 0 && index-5 < len(PLOT_LAYOUT) { // ABOVE
			checkPlots = append(checkPlots, PLOT_LAYOUT[index-5])
		} else if index+1 >= 0 && index+1 < len(PLOT_LAYOUT) { // RIGHT
			checkPlots = append(checkPlots, PLOT_LAYOUT[index+1])
		} else if index+5 >= 0 && index+5 < len(PLOT_LAYOUT) { // below
			checkPlots = append(checkPlots, PLOT_LAYOUT[index+5])
		} else if index-1 >= 0 && index-1 < len(PLOT_LAYOUT) { // left
			checkPlots = append(checkPlots, PLOT_LAYOUT[index-1])
		}

		hasAdjacentUnlocked := false
		for _, plotId := range checkPlots {
			if slices.Contains(gardenData.UnlockedPlotsIds, plotId) {
				hasAdjacentUnlocked = true
				break
			}
		}

		// BARN SKIN
		if index == 12 {
			item := notenoughupdates.NEUConstants.Garden.BarnSkins[gardenData.SelectedBarnSkin]
			if item == nil {
				item = notenoughupdates.NEUConstants.Garden.BarnSkins["default_1"]
				output.BarnSkin = utility.TitleCase(gardenData.SelectedBarnSkin)
			} else {
				output.BarnSkin = utility.GetRawLore(item.Name)
			}

			output.Layout = append(output.Layout, models.ProcessedItem{
				DisplayName: item.Name,
				Texture:     fmt.Sprintf("http://localhost:8080/api/item/%s", strings.ReplaceAll(item.ItemId, "-", ":")),
			})
		}

		textureId := "STAINED_GLASS_PANE:14"
		if slices.Contains(gardenData.UnlockedPlotsIds, plot) {
			textureId = "GRASS"
		} else if hasAdjacentUnlocked {
			textureId = "WOOD_BUTTON"
		}

		output.Layout = append(output.Layout, models.ProcessedItem{
			DisplayName: PLOT_NAMES[plot],
			Texture:     fmt.Sprintf("http://localhost:8080/api/item/%s", textureId),
		})

	}

	return output
}

func GetGarden(gardenData *skycrypttypes.Garden) *Garden {
	return &Garden{
		Level:          stats.GetLevelByXp(int(gardenData.Experience), &stats.ExtraSkillData{Type: "garden"}),
		Visitors:       getVisitors(gardenData),
		CropMilestones: getCropMilestones(gardenData),
		CropUpgrades:   getCropUpgrades(gardenData),
		Composter:      getComposter(gardenData),
		Plot:           getPlotLayout(gardenData),
	}
}
