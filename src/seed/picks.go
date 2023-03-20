package seed

import "github.com/nuck7/nfl-picks-server/src/models"

var Picks = []models.Pick{
	{
		MatchupID: 1,
		UserID:    1,
		WinnerID:  1,
	},
	{
		MatchupID: 2,
		UserID:    1,
		WinnerID:  4,
	},
	{
		MatchupID: 3,
		UserID:    1,
		WinnerID:  5,
	},
	{
		MatchupID: 4,
		UserID:    1,
		WinnerID:  7,
	},
	{
		MatchupID: 5,
		UserID:    1,
		WinnerID:  10,
	},
}
