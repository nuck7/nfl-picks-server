package types

import "time"

type CreateUserInput struct {
	Name  string
	Email string
}

type CreateTeamInput struct {
	Name string
	City string
}

type CreateMatchupInput struct {
	HomeTeamID string
	AwayTeamID string
	WeekID     string
}

type CreateWeekInput struct {
	Name  string
	Start string
	End   string
}

type CreatePickInput struct {
	UserID    string
	MatchupID string
	WinnerID  string
}

type MatchupResponse struct {
	ID           uint
	HomeTeamID   uint
	HomeTeamCity string
	HomeTeamName string
	AwayTeamID   uint
	AwayTeamCity string
	AwayTeamName string
	WeekID       uint
}

type GetWeekResponse struct {
	ID        uint
	Name      string
	Start     time.Time
	End       time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Matchups  []MatchupResponse
}

type WeekMatchupView struct {
	Week_id        uint
	Week_name      string
	Week_start     time.Time
	Week_end       time.Time
	Matchup_id     uint
	Home_team_id   uint
	Home_team_city string
	Home_team_name string
	Away_team_id   uint
	Away_team_city string
	Away_team_name string
}

type WeekMatchupResponse struct {
	ID       uint
	Name     string
	Start    time.Time
	End      time.Time
	Matchups []MatchupResponse
}
