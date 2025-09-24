package models

import (
	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

type HypixelGardenResponse struct {
	Success bool                 `json:"success"`
	Cause   string               `json:"cause,omitempty"`
	Garden  skycrypttypes.Garden `json:"garden"`
}

type Garden struct {
	Level          Skill           `json:"level"`
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
	Name    string `json:"name"`
	Texture string `json:"texture"`
	Level   Skill  `json:"level"`
}

type CropUpgrade struct {
	Name    string `json:"name"`
	Texture string `json:"texture"`
	Level   Skill  `json:"level"`
}

type PlotLayout struct {
	Unlocked int             `json:"unlocked"`
	Total    int             `json:"total"`
	BarnSkin string          `json:"barnSkin"`
	Layout   []ProcessedItem `json:"layout"`
}
