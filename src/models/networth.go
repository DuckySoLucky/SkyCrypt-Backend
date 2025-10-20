package models

type NetworthResult struct {
	Networth            float64                  `json:"networth"`
	UnsoulboundNetworth float64                  `json:"unsoulboundNetworth"`
	NoInventory         bool                     `json:"noInventory"`
	IsNonCosmetic       bool                     `json:"isNonCosmetic"`
	Purse               float64                  `json:"purse"`
	Bank                float64                  `json:"bank"`
	PersonalBank        float64                  `json:"personalBank"`
	Types               map[string]*NetworthType `json:"types"`
}

type NetworthType struct {
	Total            float64 `json:"total"`
	UnsoulboundTotal float64 `json:"unsoulboundTotal"`
}

type Networth struct {
	Networth            NetworthResult `json:"networth"`
	NonCosmeticNetworth NetworthResult `json:"nonCosmeticNetworth"`
}
