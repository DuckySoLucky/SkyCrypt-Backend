package stats

import (
	"skycrypt/src/models"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

type StatsOutput struct {
	Username        string                         `json:"username"`
	DisplayName     string                         `json:"displayName"`
	UUID            string                         `json:"uuid"`
	ProfileID       string                         `json:"profile_id"`
	ProfileCuteName string                         `json:"profile_cute_name"`
	Selected        bool                           `json:"selected"`
	Profiles        []*models.ProfilesStats        `json:"profiles"`
	Members         []*models.MemberStats          `json:"members"`
	Social          skycrypttypes.SocialMediaLinks `json:"social"`
	Rank            *models.RankOutput             `json:"rank"`
	Skills          *models.Skills                 `json:"skills"`
	SkyBlockLevel   models.Skill                   `json:"skyblock_level"`
	Joined          int64                          `json:"joined"`
	Purse           float64                        `json:"purse"`
	Bank            *float64                       `json:"bank"`
	PersonalBank    float64                        `json:"personalBank"`
	FairySouls      *models.FairySouls             `json:"fairySouls"`
	APISettings     map[string]bool                `json:"apiSettings"`
}

func GetStats(mowojang *models.MowojangReponse, profiles *models.HypixelProfilesResponse, profile *skycrypttypes.Profile, player *skycrypttypes.Player, userProfile *skycrypttypes.Member, museum *skycrypttypes.Museum, members []*models.MemberStats) (*StatsOutput, error) {
	return &StatsOutput{
		Username:        mowojang.Name,
		DisplayName:     mowojang.Name,
		UUID:            mowojang.UUID,
		ProfileID:       profile.ProfileID,
		ProfileCuteName: profile.CuteName,
		Selected:        profile.Selected,
		Profiles:        FormatProfiles(profiles),
		Members:         members,
		Social:          player.SocialMedia.Links,
		Rank:            GetRank(player),
		Skills:          GetSkills(userProfile, profile, player),
		SkyBlockLevel:   GetSkyBlockLevel(userProfile),
		Joined:          userProfile.Profile.FirstJoin,
		Purse:           userProfile.Currencies.CoinPurse,
		Bank:            profile.Banking.Balance,
		PersonalBank:    userProfile.Profile.BankAccount,
		FairySouls:      GetFairySouls(userProfile, profile.GameMode),
		APISettings:     GetAPISettings(userProfile, profile, museum),
	}, nil
}
