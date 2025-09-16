package stats

import (
	"fmt"
	notenoughupdates "skycrypt/src/NotEnoughUpdates"

	"skycrypt/src/constants"
	"skycrypt/src/lib"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"slices"
	"strings"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

func ProcessItems(items *[]skycrypttypes.Item, source string) []models.ProcessedItem {
	var processedItems []models.ProcessedItem
	for _, item := range *items {
		processedItem := ProcessItem(&item, source)
		processedItems = append(processedItems, processedItem)
	}

	return processedItems
}

func ProcessItem(item *skycrypttypes.Item, source string) models.ProcessedItem {
	if item.Tag == nil {
		return models.ProcessedItem{}
	}

	processedItem := models.ProcessedItem{
		Item:        *item,
		DisplayName: item.Tag.Display.Name,
		Lore:        item.Tag.Display.Lore,
		Source:      source,
	}

	rawLore := make([]string, len(processedItem.Lore))
	for i, lore := range processedItem.Lore {
		rawLore[i] = utility.GetRawLore(lore)
	}

	itemType := ParseItemTypeFromLore(rawLore, *item)
	processedItem.Rarity = itemType.Rarity
	processedItem.Categories = itemType.Categories
	if processedItem.Recombobulated {
		processedItem.Lore = append(processedItem.Lore, "§8(Recombobulated)")
	}

	if item.Tag.Display.Color != 0 {
		color := fmt.Sprintf("#%06X", item.Tag.Display.Color)
		if item.Tag.ExtraAttributes.DyeItem != "" {
			defaultHexColor := constants.ITEMS[item.Tag.ExtraAttributes.Id].Color
			if defaultHexColor != "" {
				fmt.Printf("[CUSTOM_RESOURCES] Using default color for item %s: %s\n", item.Tag.ExtraAttributes.Id, defaultHexColor)
				color = defaultHexColor
			}
		}

		if !slices.Contains(constants.BLACKLISTED_HEX_ARMOR_PIECES, item.Tag.ExtraAttributes.Id) {
			processedItem.Lore = append(processedItem.Lore, "", fmt.Sprintf("§7Color: %s", color))
		}
	}

	if item.Tag.ExtraAttributes != nil {
		processedItem.Recombobulated = item.Tag.ExtraAttributes.Recombobulated == 1
		if item.Tag.SkullOwner == nil {
			// Do not apply shiny effecet to skulls
			processedItem.Shiny = len(item.Tag.ExtraAttributes.Enchantments) > 0
		}

		// Timestamps
		if item.Tag.ExtraAttributes.Timestamp != nil {
			if timestamp, ok := item.Tag.ExtraAttributes.Timestamp.(float64); ok {
				processedItem.Lore = append(processedItem.Lore, "", fmt.Sprintf("§7Obtained: §c{TIMESTAMP:%.0f}", timestamp))
			} else if timestamp, ok := item.Tag.ExtraAttributes.Timestamp.(string); ok {
				parsedTimestamp := utility.ParseTimestamp(timestamp)
				processedItem.Lore = append(processedItem.Lore, "", fmt.Sprintf("§7Obtained: §c{TIMESTAMP:%d}", parsedTimestamp))
			} else if timestamp, ok := item.Tag.ExtraAttributes.Timestamp.(int64); ok {
				processedItem.Lore = append(processedItem.Lore, "", fmt.Sprintf("§7Obtained: §c{TIMESTAMP:%d}", timestamp))
			} else {
				fmt.Printf("Unexpected type for timestamp: %T, %s\n", item.Tag.ExtraAttributes.Timestamp, item.Tag.ExtraAttributes.Timestamp)
			}
		}

		// Gemstones
		if item.Tag.ExtraAttributes.Gems != nil {
			gems := ParseItemGems(item.Tag.ExtraAttributes.Gems, itemType.Rarity)
			if len(gems) > 0 {
				processedItem.Lore = append(processedItem.Lore, "", "§7Applied Gemstones:")
				for _, gem := range gems {
					processedItem.Lore = append(processedItem.Lore, fmt.Sprintf("§7 - %s", gem.Lore))
				}
			}
		}

		// Levelable enchantments
		if item.Tag.ExtraAttributes.HecatombSRuns != 0 {
			AddLevelableEnchantmentsToLore(item.Tag.ExtraAttributes.HecatombSRuns, constants.ENCHANTMENT_LADDERS["hecatomb_s_runs"], &processedItem.Lore)
		}

		if item.Tag.ExtraAttributes.ChampionCombatXP != 0 {
			AddLevelableEnchantmentsToLore(int(item.Tag.ExtraAttributes.ChampionCombatXP), constants.ENCHANTMENT_LADDERS["champion_combat_xp"], &processedItem.Lore)
		}

		if item.Tag.ExtraAttributes.FarmedCultivating != 0 {
			AddLevelableEnchantmentsToLore(item.Tag.ExtraAttributes.FarmedCultivating, constants.ENCHANTMENT_LADDERS["farmed_cultivating"], &processedItem.Lore)
		}

		if item.Tag.ExtraAttributes.ExpertiseKills != 0 {
			AddLevelableEnchantmentsToLore(item.Tag.ExtraAttributes.ExpertiseKills, constants.ENCHANTMENT_LADDERS["expertise_kills"], &processedItem.Lore)
		}

		if item.Tag.ExtraAttributes.CompactBlocks != 0 {
			AddLevelableEnchantmentsToLore(item.Tag.ExtraAttributes.CompactBlocks, constants.ENCHANTMENT_LADDERS["compact_blocks"], &processedItem.Lore)
		}

		// Wiki links
		NEUItem, err := notenoughupdates.GetItem(item.Tag.ExtraAttributes.Id)
		if err == nil && len(NEUItem.Wiki) > 0 {
			processedItem.Wiki = &models.WikipediaLinks{}
			if len(NEUItem.Wiki) == 1 {
				if strings.HasPrefix(NEUItem.Wiki[0], "https://wiki.hypixel.net/") {
					processedItem.Wiki.Official = NEUItem.Wiki[0]
				} else {
					processedItem.Wiki.Fandom = NEUItem.Wiki[0]
				}
			} else {
				if strings.HasPrefix(NEUItem.Wiki[0], "https://wiki.hypixel.net/") {
					processedItem.Wiki.Official = NEUItem.Wiki[0]
					processedItem.Wiki.Fandom = NEUItem.Wiki[1]
				} else {
					processedItem.Wiki.Fandom = NEUItem.Wiki[0]
					processedItem.Wiki.Official = NEUItem.Wiki[1]
				}
			}
		}
	}

	// POTIONS
	if *item.ID == 373 {
		color := constants.POTION_COLORS[*item.Damage]
		var potionType string
		if *item.Damage&16384 != 0 {
			potionType = "splash"
		} else {
			potionType = "normal"
		}

		processedItem.Texture = "http://localhost:8080/api/potion/" + potionType + "/" + color
	}

	if processedItem.Texture == "" {
		TextureItem := models.TextureItem{
			Count:  item.Count,
			Damage: item.Damage,
			ID:     item.ID,
			Tag:    item.Tag.ToMap(),
		}

		processedItem.Texture = lib.ApplyTexture(TextureItem)
		if strings.HasPrefix(processedItem.Texture, "http://localhost:8080/assets/resourcepacks/FurfSky/") {
			processedItem.TexturePack = "FURFSKY_REBORN"
		}

		if processedItem.Texture == "" {
			fmt.Printf("[CUSTOM_RESOURCES] Found no textures for item: %s\n", item.Tag.ExtraAttributes.Id)
		}
	}

	if item.ContainsItems != nil {
		processedItem.ContainsItems = ProcessItems(&item.ContainsItems, source)
	}

	/*if item.Tag.ExtraAttributes.ID != "" {
		prices, err := skyhelpernetworthgo.GetPrices(true, 69420, 1)
		if err == nil {
			itemCalculator, err := skyhelpernetworthgo.CalculateItem(item, prices, nil)
			if err == nil {
				processedItem.Lore = append(processedItem.Lore, fmt.Sprintf("§BALLS: %s", utility.FormatNumber(itemCalculator.Price)))
			} else {I

				fmt.Printf("You fucked up %v\n", err)
			}
		}
	}*/

	// TODO: add cake bag & legacy backpack support

	return processedItem
}
