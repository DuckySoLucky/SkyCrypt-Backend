package models

type EnchantingOutput struct {
	Data     map[string]EnchantingGameData `json:"data"`
	Unlocked bool                          `json:"unlocked"`
}

type EnchantingGameData struct {
	Name  string              `json:"name"`
	Stats EnchantingGameStats `json:"stats"`
}

type EnchantingGameStats struct {
	LastAttempt int64            `json:"lastAttempt"`
	LastClaimed int64            `json:"lastClaimed"`
	BonusClicks int              `json:"bonusClicks"`
	Games       []EnchantingGame `json:"games"`
}

type EnchantingGame struct {
	Name      string `json:"name"`
	Texture   string `json:"texture"`
	Attempts  int    `json:"attempts"`
	Claims    int    `json:"claims"`
	BestScore int    `json:"bestScore"`
}
