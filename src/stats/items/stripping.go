package stats

import (
	"skycrypt/src/models"
)

func StripItems(items *[]models.ProcessedItem, search ...bool) []models.StrippedItem {
	output := make([]models.StrippedItem, len(*items))
	for i, item := range *items {
		output[i] = *StripItem(&item, search...)

		if len(item.ContainsItems) > 0 {
			output[i].ContainsItems = StripItems(&item.ContainsItems, search...)
		}
	}

	return output
}

func StripItem(item *models.ProcessedItem, search ...bool) *models.StrippedItem {
	if item == nil {
		return nil
	}

	output := &models.StrippedItem{
		DisplayName:    item.DisplayName,
		Lore:           item.Lore,
		Rarity:         item.Rarity,
		Recombobulated: item.Recombobulated,
		Texture:        item.Texture,
		ContainsItems:  make([]models.StrippedItem, len(item.ContainsItems)),
		Shiny:          item.Shiny,
	}

	if len(search) > 0 && search[0] {
		output.Source = item.Source
	}

	if item.IsInactive != nil {
		output.IsInactive = item.IsInactive
	}

	if item.Count != nil && *item.Count > 1 {
		output.Count = item.Count
	}

	if item.Wiki != nil {
		output.Wiki = item.Wiki
	}

	if item.TexturePack != "" {
		output.TexturePack = item.TexturePack
	}

	return output
}

func StripPets(pets []models.ProcessedPet) []models.StrippedPet {
	output := make([]models.StrippedPet, len(pets))
	for i, pet := range pets {
		output[i] = StripPet(pet)
	}

	return output
}

func StripPet(pet models.ProcessedPet) models.StrippedPet {
	return models.StrippedPet{
		Active:      pet.Active,
		Type:        pet.Type,
		Rarity:      pet.Rarity,
		Level:       pet.Level.Level,
		DisplayName: pet.Name,
		Texture:     pet.Texture,
		Lore:        pet.Lore,
		Stats:       pet.Stats,
	}
}
