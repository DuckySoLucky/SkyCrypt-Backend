package constants

type essence struct {
	Name    string
	Texture string
}

var ESSENCE = map[string]essence{
	"ice": {
		Name:    "Ice",
		Texture: "http://localhost:8080/api/head/ddba642efffa13ec3730eafc5914ab68115c1f998803f74452e2e0cd26af0b8",
	},
	"wither": {
		Name:    "Wither",
		Texture: "http://localhost:8080/api/head/c4db4adfa9bf48ff5d41707ae34ea78bd2371659fcd8cd8934749af4cce9b",
	},
	"spider": {
		Name:    "Spider",
		Texture: "http://localhost:8080/api/head/16617131250e578333a441fdf4a5b8c62163640a9d06cd67db89031d03accf6",
	},
	"undead": {
		Name:    "Undead",
		Texture: "http://localhost:8080/api/head/71d7c816fc8c636d7f50a93a0ba7aaeff06c96a561645e9eb1bef391655c531",
	},
	"diamond": {
		Name:    "Diamond",
		Texture: "http://localhost:8080/api/head/964f25cfff754f287a9838d8efe03998073c22df7a9d3025c425e3ed7ff52c20",
	},
	"dragon": {
		Name:    "Dragon",
		Texture: "http://localhost:8080/api/head/33ff416aa8bec1665b92701fbe68a4effff3d06ed9147454fa77712dd6079b33",
	},
	"gold": {
		Name:    "Gold",
		Texture: "http://localhost:8080/api/head/8816606260779b23ed15f87c56c932240db745f86f683d1f4deb83a4a125fa7b",
	},
	"crimson": {
		Name:    "Crimson",
		Texture: "http://localhost:8080/api/head/67c41930f8ff0f2b0430e169ae5f38e984df1244215705c6f173862844543e9d",
	},
}

var RACE_NAMES = map[string]string{
	"crystal_core":    "Crystal Core",
	"giant_mushroom":  "Giant Mushroom",
	"precursor_ruins": "Precursor Ruins",
	"foraging_race":   "Foraging",
	"end_race":        "End",
	"chicken_race_2":  "Chicken",
	"rift_race":       "Rift",
}

var MILESTONE_RARITIES = []string{"common", "uncommon", "rare", "epic", "legendary"}

var PET_MILESTONES = map[string][]int{
	"sea_creatures_killed": {250, 1000, 2500, 5000, 10000},
	"ores_mined":           {2500, 7500, 20000, 100000, 250000},
}

var PROFILE_UPGRADES = map[string]int{
	"island_size":     10,
	"minion_slots":    5,
	"guests_count":    5,
	"coop_slots":      3,
	"coins_allowance": 5,
}

var CLAIMABLE_ITEMS = map[string]string{
	"claimed_potato_talisman":       "Potato Talisman",
	"claimed_potato_basket":         "Potato Basket",
	"claim_potato_war_silver_medal": "Silver Medal (Potato War)",
	"claim_potato_war_crown":        "Crown (Potato War)",
	"skyblock_free_cookie":          "Free Booster Cookie",
}

var BANK_COOLDOWN = map[int]string{
	1: "20 minutes",
	2: "5 minutes",
	3: "None",
}
