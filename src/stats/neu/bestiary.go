package neustats

import (
	"fmt"
	"os"
	neu "skycrypt/src/models/NEU"
	"skycrypt/src/utility"
)

func GetTexture(mob neu.NEUBestiaryRawMob) string {
	if mob.Texture == "" {
		if os.Getenv("DEV") != "true" {
			return fmt.Sprintf("/api/item/%s", mob.Item)
		}

		return fmt.Sprintf(`http://localhost:8080/api/item/%s`, mob.Item)
	}

	skinHash := utility.GetSkinHash(mob.Texture)
	if skinHash == "" {
		if os.Getenv("DEV") != "true" {
			return "/api/item/BARRIER"
		}

		return "http://localhost:8080/api/item/BARRIER"

	}

	if os.Getenv("DEV") != "true" {
		return fmt.Sprintf("/api/head/%s", skinHash)
	}

	return fmt.Sprintf("http://localhost:8080/api/head/%s", skinHash)
}

func GetIslandTexture(island neu.NEUBestiaryRawIslandData) string {
	if island.Icon.Texture == "" {
		if os.Getenv("DEV") != "true" {
			return fmt.Sprintf("/api/item/%s", island.Icon.Item)
		}

		return fmt.Sprintf(`http://localhost:8080/api/item/%s`, island.Icon.Item)
	}

	if os.Getenv("DEV") != "true" {
		return fmt.Sprintf("/api/head/%s", utility.GetSkinHash(island.Icon.Texture))
	}

	return fmt.Sprintf("http://localhost:8080/api/head/%s", utility.GetSkinHash(island.Icon.Texture))
}

func FormatBestiaryConstants(bestiaryConstants neu.NEUBestiaryRaw) neu.BestiaryConstants {
	formatted := neu.BestiaryConstants{
		Brackets: bestiaryConstants.Brackets,
		Islands:  make(map[string]neu.BestiaryCategory),
	}

	for name, islandData := range bestiaryConstants.Islands {
		if islandData.HasSubcategories {
			for subcategoryName, subcategoryMobs := range islandData.SubcategoryIslands {
				category := neu.BestiaryCategory{
					Name:    utility.GetRawLore(subcategoryName),
					Texture: GetIslandTexture(islandData),
					Mobs:    make([]neu.BestiaryMob, len(subcategoryMobs.Mobs)),
				}

				for i, mob := range subcategoryMobs.Mobs {
					category.Mobs[i] = neu.BestiaryMob{
						Name:    utility.GetRawLore(mob.Name),
						Texture: GetTexture(mob),
						Cap:     mob.Cap,
						Mobs:    mob.Mobs,
						Bracket: mob.Bracket,
					}
				}

				formatted.Islands[fmt.Sprintf("%s_%s", name, subcategoryName)] = category
			}
			continue
		}

		category := neu.BestiaryCategory{
			Name:    utility.GetRawLore(islandData.Name),
			Texture: GetIslandTexture(islandData),
			Mobs:    make([]neu.BestiaryMob, len(islandData.Mobs)),
		}

		for i, mob := range islandData.Mobs {
			category.Mobs[i] = neu.BestiaryMob{
				Name:    utility.GetRawLore(mob.Name),
				Texture: GetTexture(mob),
				Cap:     mob.Cap,
				Mobs:    mob.Mobs,
				Bracket: mob.Bracket,
			}
		}

		formatted.Islands[name] = category
	}

	return formatted
}
