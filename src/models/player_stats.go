package models

type StatsInfo map[string]int

type Stats struct {
	Stats map[string]StatsInfo `json:"stats"`
}
