package models

import skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"

type HypixelProfilesResponse struct {
	Success  bool                    `json:"success"`
	Cause    string                  `json:"cause,omitempty"`
	Profiles []skycrypttypes.Profile `json:"profiles"`
}
