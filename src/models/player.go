package models

import skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"

type HypixelPlayerResponse struct {
	Success bool                 `json:"success"`
	Cause   string               `json:"cause,omitempty"`
	Player  skycrypttypes.Player `json:"player"`
}
