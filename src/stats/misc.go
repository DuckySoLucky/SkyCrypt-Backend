package stats

import (
	"fmt"
	"os"
	"skycrypt/src/constants"
	"skycrypt/src/utility"
	"strings"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

type MiscOutput struct {
	Essence           []MiscEssence               `json:"essence"`
	Kills             MiscKills                   `json:"kills"`
	Gifts             MiscGifts                   `json:"gifts"`
	SeasonOfJerry     MiscSeasonOfJerry           `json:"season_of_jerry"`
	Dragons           MiscDragons                 `json:"dragons"`
	EndstoneProtector MiscEndstoneProtector       `json:"endstone_protector"`
	Damage            MiscDamage                  `json:"damage"`
	PetMilestones     map[string]MiscPetMilestone `json:"pet_milestones"`
	MythologicalEvent MiscMythologicalEvent       `json:"mythological_event"`
	ProfileUpgrades   MiscProfileUpgrades         `json:"profile_upgrades"`
	Auctions          MiscAuctions                `json:"auctions"`
	ClaimedItems      map[string]int64            `json:"claimed_items"`
	Uncategorized     map[string]any              `json:"uncategorized"`
}

type MiscAuctions struct {
	Bids        float64            `json:"bids"`
	HighestBid  float64            `json:"highest_bid"`
	Won         float64            `json:"won"`
	TotalBought map[string]float64 `json:"total_bought"`
	GoldSpent   float64            `json:"gold_spent"`
	Created     float64            `json:"created"`
	Fees        float64            `json:"fees"`
	TotalSold   map[string]float64 `json:"total_sold"`
	GoldEarned  float64            `json:"gold_earned"`
	NoBids      float64            `json:"no_bids"`
}

type MiscProfileUpgrades map[string]int

type MiscMythologicalEvent struct {
	Kills                 float64            `json:"kills"`
	BurrowsDugNext        map[string]float64 `json:"burrows_dug_next"`
	BurrowsDugCombat      map[string]float64 `json:"burrows_dug_combat"`
	BurrowsDugTreasure    map[string]float64 `json:"burrows_dug_treasure"`
	BurrowsChainsComplete map[string]float64 `json:"burrows_chains_complete"`
}

type MiscPetMilestone struct {
	Amount   int    `json:"amount"`
	Rarity   string `json:"rarity"`
	Total    int    `json:"total"`
	Progress string `json:"progress"`
}

type MiscDamage struct {
	HighestCriticalDamage float64 `json:"highest_critical_damage"`
}

type MiscEndstoneProtector struct {
	Kills  int `json:"kills"`
	Deaths int `json:"deaths"`
}

type MiscDragons struct {
	EnderCrystalsDestroyed int                `json:"ender_crystals_destroyed"`
	MostDamage             map[string]float64 `json:"most_damage"`
	FastestKill            map[string]float64 `json:"fastest_kill"`
	LastHits               map[string]float64 `json:"last_hits"`
	Deaths                 map[string]float64 `json:"deaths"`
}

type MiscSeasonOfJerry struct {
	MostSnowballsHit     int `json:"most_snowballs_hit"`
	MostDamageDealt      int `json:"most_damage_dealt"`
	MostMagmaDamageDealt int `json:"most_magma_damage_dealt"`
	MostCannonballsHit   int `json:"most_cannonballs_hit"`
}

type MiscEssence struct {
	Name    string `json:"name"`
	Texture string `json:"texture"`
	Amount  int    `json:"amount"`
}

type MiscKills struct {
	TotalKills  int        `json:"total_kills"`
	TotalDeaths int        `json:"total_deaths"`
	Kills       []MiscKill `json:"kills"`
	Deaths      []MiscKill `json:"deaths"`
}

type MiscKill struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type MiscGifts struct {
	Given    int `json:"given"`
	Received int `json:"received"`
}

func getEssence(userProfile *skycrypttypes.Member) []MiscEssence {
	essence := make([]MiscEssence, 0, len(constants.ESSENCE))
	for essenceId, essenceData := range constants.ESSENCE {
		essenceData := MiscEssence{
			Name:    essenceData.Name,
			Texture: essenceData.Texture,
			Amount:  userProfile.Currencies.Essence[strings.ToUpper(essenceId)].Current,
		}

		if os.Getenv("DEV") == "true" {
			essenceData.Texture = strings.Replace(essenceData.Texture, "/api/head/", "http://localhost:8080/api/head/", 1)
		}

		essence = append(essence, essenceData)
	}

	return essence
}

func getKills(userProfile *skycrypttypes.Member) MiscKills {
	totalKills, totalDeaths := 0, 0
	kills, deaths := []MiscKill{}, []MiscKill{}
	for id, amount := range userProfile.PlayerStats.Kills {
		if id == "total" {
			continue
		}

		name := constants.MOB_NAMES[id]
		if name == "" {
			name = utility.TitleCase(id)
		}

		totalKills += int(amount)
		kills = append(kills, MiscKill{
			Name:   name,
			Amount: int(amount),
		})
	}

	for id, amount := range userProfile.PlayerStats.Deaths {
		if id == "total" {
			continue
		}

		totalDeaths += int(amount)
		name := constants.MOB_NAMES[id]
		if name == "" {
			name = utility.TitleCase(id)
		}

		deaths = append(deaths, MiscKill{
			Name:   name,
			Amount: int(amount),
		})
	}

	return MiscKills{
		TotalKills:  totalKills,
		TotalDeaths: totalDeaths,
		Kills:       kills,
		Deaths:      deaths,
	}
}

func getGifts(userProfile *skycrypttypes.Member) MiscGifts {
	return MiscGifts{
		Given:    int(userProfile.PlayerStats.Gifts.Given),
		Received: int(userProfile.PlayerStats.Gifts.Received),
	}
}

func getSeasonOfJerry(userProfile *skycrypttypes.Member) MiscSeasonOfJerry {
	return MiscSeasonOfJerry{
		MostSnowballsHit:     int(userProfile.PlayerStats.WinterIslandData.MostSnowballsHit),
		MostDamageDealt:      int(userProfile.PlayerStats.WinterIslandData.MostDamageDealt),
		MostMagmaDamageDealt: int(userProfile.PlayerStats.WinterIslandData.MostMagmaDamageDealt),
		MostCannonballsHit:   int(userProfile.PlayerStats.WinterIslandData.MostCannonballsHit),
	}
}

func getDragons(userProfile *skycrypttypes.Member) MiscDragons {
	dragonKills, dragonKillsTotal, dragonDeaths, dragonDeathsTotal := map[string]float64{}, 0.0, map[string]float64{}, 0.0
	for mobId, amount := range userProfile.PlayerStats.Kills {
		if strings.HasPrefix(mobId, "master_wither_king") {
			continue
		}

		if strings.HasSuffix(mobId, "_dragon") {
			dragonId := strings.ReplaceAll(mobId, "_dragon", "")
			dragonKills[dragonId] += float64(amount)
			dragonKillsTotal += float64(amount)
		}
	}

	dragonKills["total"] = dragonKillsTotal

	for mobId, amount := range userProfile.PlayerStats.Deaths {
		if strings.HasPrefix(mobId, "master_wither_king") {
			continue
		}

		if strings.HasSuffix(mobId, "_dragon") {
			dragonId := strings.ReplaceAll(mobId, "_dragon", "")
			dragonDeaths[dragonId] += float64(amount)
			dragonDeathsTotal += float64(amount)
		}
	}

	dragonDeaths["total"] = dragonDeathsTotal

	return MiscDragons{
		EnderCrystalsDestroyed: int(userProfile.PlayerStats.EndIsland.DragonFight.EnderCrystalsDestroyed),
		MostDamage:             userProfile.PlayerStats.EndIsland.DragonFight.MostDamage,
		FastestKill:            userProfile.PlayerStats.EndIsland.DragonFight.FastestKill,
		LastHits:               dragonKills,
		Deaths:                 dragonDeaths,
	}
}

func getEndstoneProtector(userProfile *skycrypttypes.Member) MiscEndstoneProtector {
	return MiscEndstoneProtector{
		Kills:  int(userProfile.PlayerStats.Kills["corrupted_protector"]),
		Deaths: int(userProfile.PlayerStats.Deaths["corrupted_protector"]),
	}
}

func getDamage(userProfile *skycrypttypes.Member) MiscDamage {
	return MiscDamage{
		HighestCriticalDamage: userProfile.PlayerStats.HighestCriticalDamage,
	}
}

func getPetMilestone(typeName string, amount float64) MiscPetMilestone {
	rarity := "common"
	milestones := constants.PET_MILESTONES[typeName]
	lastIndex := -1
	for i := len(milestones) - 1; i >= 0; i-- {
		if amount >= float64(milestones[i]) {
			lastIndex = i
			break
		}
	}
	if lastIndex >= 0 && lastIndex < len(constants.MILESTONE_RARITIES) {
		rarity = constants.MILESTONE_RARITIES[lastIndex]
	}
	total := int(amount)
	progress := "0"
	if amount > 0 && len(milestones) > 0 {
		maxMilestone := float64(milestones[len(milestones)-1])
		if maxMilestone > 0 {
			p := (amount / maxMilestone) * 100
			if p > 100 {
				p = 100
			}

			progress = fmt.Sprintf("%.2f%%", p)
		}
	}
	return MiscPetMilestone{
		Amount:   int(amount),
		Rarity:   rarity,
		Total:    total,
		Progress: progress,
	}
}

func getPetMilestones(userProfile *skycrypttypes.Member) map[string]MiscPetMilestone {
	return map[string]MiscPetMilestone{
		"sea_creatures_killed": getPetMilestone("sea_creatures_killed", userProfile.PlayerStats.Pets.Milestone.SeaCreaturesKilled),
		"ores_mined":           getPetMilestone("ores_mined", userProfile.PlayerStats.Pets.Milestone.OresMined),
	}
}

func getMythologicalEvent(userProfile *skycrypttypes.Member) MiscMythologicalEvent {
	return MiscMythologicalEvent{
		Kills:                 userProfile.PlayerStats.Mythos.Kills,
		BurrowsDugNext:        userProfile.PlayerStats.Mythos.BurrowsDugNext,
		BurrowsDugCombat:      userProfile.PlayerStats.Mythos.BurrowsDugCombat,
		BurrowsDugTreasure:    userProfile.PlayerStats.Mythos.BurrowsDugTreasure,
		BurrowsChainsComplete: userProfile.PlayerStats.Mythos.BurrowsChainsComplete,
	}
}

func getProfileUpgrades(profile *skycrypttypes.Profile) MiscProfileUpgrades {
	output := MiscProfileUpgrades{}
	for upgrade := range constants.PROFILE_UPGRADES {
		output[upgrade] = 0
	}

	if profile.CommunityUpgrades != nil && profile.CommunityUpgrades.UpgradeStates != nil {
		for _, u := range profile.CommunityUpgrades.UpgradeStates {
			if u.Tier > output[u.Upgrade] {
				output[u.Upgrade] = u.Tier
			}
		}
	}

	return output
}

func getAuctions(userProfile *skycrypttypes.Member) MiscAuctions {
	auctions := userProfile.PlayerStats.Auctions

	totalSold, totalSoldAmount, totalBought, totalBoughtAmount := map[string]float64{}, 0.0, map[string]float64{}, 0.0
	for item, amount := range auctions.TotalSold {
		totalSold[item] = amount
		totalSoldAmount += amount
	}

	totalSold["total"] = totalSoldAmount

	for item, amount := range auctions.TotalBought {
		totalBought[item] = amount
		totalBoughtAmount += amount
	}

	totalBought["total"] = totalBoughtAmount

	return MiscAuctions{
		Bids:        auctions.Bids,
		HighestBid:  auctions.HighestBid,
		Won:         auctions.Won,
		TotalBought: totalBought,
		GoldSpent:   auctions.GoldSpent,
		Created:     auctions.Created,
		Fees:        auctions.Fees,
		TotalSold:   totalSold,
		GoldEarned:  auctions.GoldEarned,
		NoBids:      auctions.NoBids,
	}
}

func getUncategorized(userProfile *skycrypttypes.Member) map[string]any {
	personalBank := constants.BANK_COOLDOWN[userProfile.Profile.PersonalBankUpgrade]
	if personalBank == "" {
		personalBank = "Unknown"
	}

	return map[string]any{
		"soulflow":                 userProfile.ItemData.Soulflow,
		"teleporter_pill_consumed": userProfile.ItemData.TeleporterPillConsumed,
		"personal_bank":            personalBank,
		"metaphysical_serum":       userProfile.Experimentation.SerumsDrank,
		"reaper_peppers_eaten":     userProfile.PlayerData.ReaperPeppersEaten,
		"mcgrubber_burger":         userProfile.Rift.Castle.GrubberStacks,
		"wriggling_larva":          userProfile.Garden.LarvaConsumed,
		"refined_bottle_of_jyrre":  userProfile.WinterPlayerData.RefinedJyrreUses,
	}
}

func getClaimedItems(player *skycrypttypes.Player) map[string]int64 {
	return map[string]int64{
		"potato_talisman":         player.ClaimedPotatoTalisman,
		"potato_basket":           player.ClaimedPotatoBasket,
		"potato_war_silver_medal": player.ClaimPotatoWarSilverMedal,
		"potato_war_crown":        player.ClaimPotatoWarCrown,
		"skyblock_free_cookie":    player.SkyblockFreeCookie,
		"century_cake":            player.ClaimedCenturyCake,
		"century_cake_(year_200)": player.ClaimedCenturyCake200,
	}
}

func GetMisc(userProfile *skycrypttypes.Member, profile *skycrypttypes.Profile, player *skycrypttypes.Player) *MiscOutput {
	return &MiscOutput{
		Essence:           getEssence(userProfile),
		Kills:             getKills(userProfile),
		Gifts:             getGifts(userProfile),
		SeasonOfJerry:     getSeasonOfJerry(userProfile),
		Dragons:           getDragons(userProfile),
		EndstoneProtector: getEndstoneProtector(userProfile),
		Damage:            getDamage(userProfile),
		PetMilestones:     getPetMilestones(userProfile),
		MythologicalEvent: getMythologicalEvent(userProfile),
		ProfileUpgrades:   getProfileUpgrades(profile),
		Auctions:          getAuctions(userProfile),
		Uncategorized:     getUncategorized(userProfile),
		ClaimedItems:      getClaimedItems(player),
	}
}
