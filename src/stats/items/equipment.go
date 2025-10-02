package stats

import (
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"slices"
)

func GetEquipment(equipment []models.ProcessedItem) models.EquipmentResult {
	if utility.Every(equipment, isInvalidItem) {
		return models.EquipmentResult{
			Equipment: []models.StrippedItem{},
			Stats:     map[string]float64{},
		}
	}

	validItems := utility.Filter(equipment, func(item models.ProcessedItem) bool {
		return !isInvalidItem(item)
	})
	if len(validItems) == 0 {
		return models.EquipmentResult{
			Equipment: []models.StrippedItem{},
			Stats:     map[string]float64{},
		}
	}

	reversedEquipment := make([]models.ProcessedItem, len(validItems))
	copy(reversedEquipment, validItems)
	slices.Reverse(reversedEquipment)

	return models.EquipmentResult{
		Equipment: StripItems(&reversedEquipment),
		Stats:     GetStatsFromItems(validItems),
	}
}
