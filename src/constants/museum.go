package constants

import (
	"skycrypt/src/models"
	"slices"
	"sort"
	"strings"
	"time"
)

type MuseumConstants struct {
	ArmorSetToId map[string]string
	ArmorSets    map[string][]string
	Children     map[string]string
	Weapons      []string
	Armor        []string
	Rarities     []string
}

var MUSEUM = MuseumConstants{}

func (m *MuseumConstants) GetAllItems() []string {
	items := make([]string, 0, len(m.Weapons)+len(m.Armor)+len(m.Rarities))
	items = append(items, m.Weapons...)
	items = append(items, m.Armor...)
	items = append(items, m.Rarities...)
	return items
}

var priorityOrder = []string{"HAT", "HOOD", "HELMET", "CHESTPLATE", "TUNIC", "LEGGINGS", "TROUSERS", "SLIPPERS", "BOOTS", "NECKLACE", "CLOAK", "BELT", "GAUNTLET", "GLOVES"}

func sortMuseumItems(items []string) {
	sort.Slice(items, func(i, j int) bool {
		a := items[i]
		b := items[j]

		aItem, aOk := ITEMS[a]
		bItem, bOk := ITEMS[b]
		if !aOk || !bOk {
			return false
		}

		aId := aItem.SkyblockID
		bId := bItem.SkyblockID

		aIdx := len(priorityOrder)
		bIdx := len(priorityOrder)

		for idx, keyword := range priorityOrder {
			if strings.Contains(aId, keyword) && aIdx == len(priorityOrder) {
				aIdx = idx
			}

			if strings.Contains(bId, keyword) && bIdx == len(priorityOrder) {
				bIdx = idx
			}
		}

		return aIdx < bIdx
	})

}

func getMuseumItems() {
	output := MuseumConstants{
		Children:     make(map[string]string),
		ArmorSets:    make(map[string][]string),
		ArmorSetToId: make(map[string]string),
	}

	for _, item := range ITEMS {
		if item.MuseumData == nil {
			continue
		}

		category := strings.ToLower(item.MuseumData.Type)
		switch category {
		case "weapons":
			output.Weapons = append(output.Weapons, item.SkyblockID)
		case "rarities":
			output.Rarities = append(output.Rarities, item.SkyblockID)
		}

		if item.MuseumData.Parent != nil {
			for parentKey, parentValue := range item.MuseumData.Parent {
				output.Children[parentValue] = parentKey
			}
		}

		if item.MuseumData.ArmorSetExperience != nil {
			var armorSetId string
			for setId := range item.MuseumData.ArmorSetExperience {
				armorSetId = setId
			}

			output.ArmorSets[armorSetId] = append(output.ArmorSets[armorSetId], item.SkyblockID)

			sortMuseumItems(MUSEUM.ArmorSets[armorSetId])

			output.ArmorSetToId[armorSetId] = output.ArmorSets[armorSetId][0]

			if !slices.Contains(output.Armor, armorSetId) {
				output.Armor = append(output.Armor, armorSetId)
			}
		}
	}

	MUSEUM = output
}

func init() {
	go func() {
		getMuseumItems()
		for len(MUSEUM.Weapons) == 0 {
			time.Sleep(1 * time.Second)
			getMuseumItems()
		}

		ticker := time.NewTicker(60 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			getMuseumItems()
		}
	}()
}

var MUSEUM_INVENTORY = []models.MuseumInventoryItem{
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Museum",
			Rarity:      "rare",
			Texture:     "/api/head/438cf3f8e54afc3b3f91d20a49f324dca1486007fe545399055524c17941f4dc",
			Lore: []string{
				"§7The §9Museum §7is a compendium",
				"§7of all of your items in",
				"§7SkyBlock. Donate items to your",
				"§7Museum to unlock rewards.",
				"",
				"§7Other players can visit your",
				"§7Museum at any time! Display your",
				"§7best items proudly for all to", "§7see.",
				"",
			},
		},
		Position:     4,
		ProgressType: "total",
	},
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Weapons",
			Rarity:      "uncommon",
			Texture:     "/api/item/DIAMOND_SWORD",
			Lore: []string{
				"§7View all of the §6Weapons §7that",
				"§7you have donated to the",
				"§7§9Museum§7!",
				"",
			},
		},
		InventoryType: "weapons",
		Position:      19,
		ProgressType:  "weapons",
		ContainsItems: []models.MuseumInventoryItem{},
	},
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Armor Sets",
			Rarity:      "uncommon",
			Texture:     "/api/item/DIAMOND_CHESTPLATE",
			Lore: []string{
				"§7View all of the §9Armor Sets",
				"§9§7that you have donated to the",
				"§7§9Museum§7!",
				"",
			},
		},
		InventoryType: "armor",
		Position:      21,
		ProgressType:  "armor",
		ContainsItems: []models.MuseumInventoryItem{},
	},
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Rarities",
			Rarity:      "uncommon",
			Texture:     "/api/head/86addbd5dedad40999473be4a7f48f6236a79a0dce971b5dbd7372014ae394d",
			Lore: []string{
				"§7View all of the §5Rarities",
				"§5§7that you have donated to the",
				"§7§9Museum§7!",
				"",
			},
		},
		InventoryType: "rarities",
		Position:      23,
		ProgressType:  "rarities",
		ContainsItems: []models.MuseumInventoryItem{},
	},
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Special Items",
			Rarity:      "uncommon",
			Texture:     "/api/item/CAKE",
			Lore: []string{
				"§7View all of the §dSpecial Items",
				"§d§7that you have donated to the",
				"§7§9Museum§7!",
				"",
				"§7These items don't count towards",
				"§7Museum progress and rewards, but",
				"§7are cool nonetheless. Items that",
				"§7are §9rare §7and §6prestigious",
				"§6§7fit into this category, and",
				"§7can be displayed in the Main",
				"§7room of the Museum.",
				"",
			},
		},
		InventoryType: "special",
		Position:      25,
		ProgressType:  "special",
		ContainsItems: []models.MuseumInventoryItem{},
	},
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Museum Appraisal",
			Rarity:      "legendary",
			Texture:     "/api/item/DIAMOND",
			Lore: []string{
				"§7§6Madame Goldsworth §7offers an",
				"§7appraisal service for Museums.",
				"§7When unlocked, she will appraise",
				"§7the value of your Museum each",
				"§7time you add or remove items.",
				"",
				"§7This service also allows you to",
				"§7appear on the §6Top Valued",
				"§6§7filter in the §9Museum",
				"§9Browser§7.",
				"",
			},
		},
		Position:     40,
		ProgressType: "appraisal",
	},
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Museum Rewards",
			Rarity:      "legendary",
			Texture:     "/api/item/GOLD_BLOCK",
			Lore: []string{
				"§7Each time you donate an item to",
				"§7your Museum, the §bCurator",
				"§b§7will reward you.",
				"",
				"§7§dSpecial Items §7do not count",
				"§7towards your Museum rewards",
				"§7progress.",
				"",
				"§7Currently, most rewards are",
				"§7§ccoming soon§7, but you can",
				"§7view them anyway.",
			},
		},
		Position: 48,
	},
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Close",
			Rarity:      "special",
			Texture:     "/api/item/BARRIER",
			Lore:        []string{},
		},
		Position: 49,
	},
	{
		ProcessedItem: models.ProcessedItem{
			DisplayName: "Museum Browser",
			Rarity:      "uncommon",
			Texture:     "/api/item/SIGN",
			Lore: []string{
				"§7View the Museums of your",
				"§7friends, top valued players, and",
				"§7more!",
			},
		},
		Position: 50,
	},
}

var MUSEUM_INVENTORY_ITEM_SLOTS = []int{10, 11, 12, 13, 14, 15, 16, 19, 20, 21, 22, 23, 24, 25, 28, 29, 30, 31, 32, 33, 34, 37, 38, 39, 40, 41, 42, 43}

var MUSEUM_INVENTORY_MISSING_ITEM_TEMPLATE = map[string]models.ProcessedItem{
	"weapons": {
		DisplayName: "Missing Weapon",
		Rarity:      "special",
		Texture:     "/api/item/INK_SACK:8",
		Lore: []string{
			"§7Click on this item in your",
			"§7inventory to add it to your",
			"§7§9Museum§7!",
		},
	},
	"armor": {
		DisplayName: "Missing Armor Set",
		Rarity:      "special",
		Texture:     "/api/item/INK_SACK:8",
		Lore: []string{
			"§7Click on an armor piece in your",
			"§7inventory that belongs to this",
			"§7armor set to donate the full set",
			"§7to your Museum.",
		},
	},
	"rarities": {
		DisplayName: "Missing Rarity",
		Rarity:      "special",
		Texture:     "/api/item/INK_SACK:8",
		Lore: []string{
			"§7Click on this item in your",
			"§7inventory to add it to your",
			"§7§9Museum§7!",
		},
	},
}

var MUSEUM_INVENTORY_HIGHER_TIER_DONATED_TEMPLATE = models.ProcessedItem{
	DisplayName: "Higher Tier Donated",
	Texture:     "/api/item/INK_SACK:10",
	Rarity:      "special",
	Lore: []string{
		"§7Donated as higher tier",
	},
}
