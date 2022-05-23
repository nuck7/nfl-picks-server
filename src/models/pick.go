package models

type Pick struct {
	ID      string `json:"id"`
	User    string `json:"user"`
	Matchup string `json:"matchup"`
	Winner  string `json:"winner"`
}
