package types

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
	Start string
	End   string
}

type CreatePickInput struct {
	UserID    string
	MatchupID string
	WinnerID  string
}
