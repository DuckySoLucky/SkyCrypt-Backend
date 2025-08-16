package stats

import (
	"skycrypt/src/constants"
	"skycrypt/src/models"
)

func getGame(gameData *models.ExperimentationGame, gameId string) []models.EnchantingGame {
	var output []models.EnchantingGame
	for index, tier := range constants.EXPERIMENTS.Tiers {
		attempts := gameData.Attempts[index]
		claims := gameData.Claims[index]
		bestScore := gameData.BestScores[index]
		if attempts == 0 && claims == 0 && bestScore == 0 {
			continue
		}

		switch gameId {
		case "numbers":
			index += 2
		case "simon":
			index = min(index+1, 5)
		}

		tier = constants.EXPERIMENTS.Tiers[index]
		output = append(output, models.EnchantingGame{
			Name:      tier.Name,
			Texture:   tier.Texture,
			Attempts:  attempts,
			Claims:    claims,
			BestScore: bestScore,
		})
	}

	return output
}

func GetEnchanting(userProfie *models.Member) models.EnchantingOutput {
	if userProfie.Experimentation.ClaimsResets == nil {
		return models.EnchantingOutput{
			Unlocked: false,
		}
	}

	output := map[string]models.EnchantingGameData{}
	games := []struct {
		key      string
		gameData *models.ExperimentationGame
	}{
		{"simon", userProfie.Experimentation.Simon},
		{"numbers", userProfie.Experimentation.Numbers},
		{"pairings", userProfie.Experimentation.Pairings},
	}

	for _, g := range games {
		output[g.key] = models.EnchantingGameData{
			Name: constants.EXPERIMENTS.Games[g.key].Name,
			Stats: models.EnchantingGameStats{
				LastAttempt: g.gameData.LastAttempt,
				LastClaimed: g.gameData.LastClaimed,
				BonusClicks: g.gameData.BonusClicks,
				Games:       getGame(g.gameData, g.key),
			},
		}
	}

	return models.EnchantingOutput{
		Unlocked: true,
		Data:     output,
	}
}
