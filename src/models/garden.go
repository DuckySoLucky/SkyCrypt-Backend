package models

import skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"

type HypixelGardenResponse struct {
	Success bool      `json:"success"`
	Cause   string    `json:"cause,omitempty"`
	Garden  skycrypttypes.Garden `json:"garden"`
}

