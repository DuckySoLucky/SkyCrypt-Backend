package models

type GetMissingAccessoresOutput struct {
	Stats               map[string]float64    `json:"stats"`
	Enrichments         map[string]int        `json:"enrichments"`
	Unique              int                   `json:"unique"`
	Total               int                   `json:"total"`
	Recombobulated      int                   `json:"recombobulated"`
	TotalRecombobulated int                   `json:"totalRecombobulated"`
	SelectedPower       string                `json:"selectedPower"`
	MagicalPower        GetMagicalPowerOutput `json:"magicalPower"`
	Accessories         []StrippedItem        `json:"accessories"`
	Missing             []StrippedItem        `json:"missing"`
	Upgrades            []StrippedItem        `json:"upgrades"`
}

type MissingOutput struct {
	Upgrades     []ProcessedItem `json:"upgrades"`
	Other        []ProcessedItem `json:"other"`
	AccessoryIds []AccessoryIds  `json:"accessoryIds"`
}

type GetMagicalPowerOutput struct {
	Total       int `json:"total"`
	Accessories int `json:"accessories"`
	Abiphone    int `json:"abiphone"`
	RiftPrism   int `json:"riftPrism"`
	Hegemony    struct {
		Rarity string `json:"rarity"`
		Amount int    `json:"amount"`
	} `json:"hegemony"`
	Rarities GetMagicalPowerRarities `json:"rarities"`
}

type GetMagicalPowerRarities map[string]struct {
	Amount       int `json:"amount"`
	MagicalPower int `json:"magicalPower"`
}

type InsertAccessory struct {
	ProcessedItem
	Id         string `json:"id"`
	Rarity     string `json:"rarity"`
	IsInactive bool   `json:"isInactive"`
}

type AccessoryIds struct {
	Id     string `json:"id"`
	Rarity string `json:"rarity"`
}

type AccessoriesOutput struct {
	Accessories  []InsertAccessory `json:"accessories"`
	AccessoryIds []AccessoryIds    `json:"accessoryIds"`
}
