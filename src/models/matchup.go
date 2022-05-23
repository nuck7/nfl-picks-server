package models

type Matchup struct {
	ID       string `json:"id"`
	HomeTeam string `json:"homeTeam"`
	AwayTeam string `json:"awayTeam"`
	Week     string `json:"week"`
}
