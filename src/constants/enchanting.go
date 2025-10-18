package constants

type game struct {
	Name string `json:"name"`
}

type tier struct {
	Name    string `json:"name"`
	Texture string `json:"texture"`
}

type experiments struct {
	Games map[string]game `json:"games"`
	Tiers []tier          `json:"tiers"`
}

var EXPERIMENTS = experiments{
	Games: map[string]game{
		"simon":    {Name: "Chronomatron"},
		"numbers":  {Name: "Ultrasequencer"},
		"pairings": {Name: "Superpairs"},
	},
	Tiers: []tier{
		{Name: "Beginner", Texture: "/api/item/INK_SACK:12"},
		{Name: "High", Texture: "/api/item/INK_SACK:10"},
		{Name: "Grand", Texture: "/api/item/INK_SACK:11"},
		{Name: "Supreme", Texture: "/api/item/INK_SACK:14"},
		{Name: "Transcendent", Texture: "/api/item/INK_SACK:1"},
		{Name: "Metaphysical", Texture: "/api/item/INK_SACK:13"},
	},
}
