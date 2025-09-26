package stats

import (
	"skycrypt/src/models"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

func GetStats(mowojang *models.MowojangReponse, profiles *models.HypixelProfilesResponse, profile *skycrypttypes.Profile, player *skycrypttypes.Player, userProfile *skycrypttypes.Member, museum *skycrypttypes.Museum, members []*models.MemberStats) (*models.StatsOutput, error) {
	return &models.StatsOutput{
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
