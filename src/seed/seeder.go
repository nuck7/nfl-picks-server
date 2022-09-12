package seed

import (
	"errors"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/nuck7/nfl-picks-server/src/models"
	"github.com/nuck7/nfl-picks-server/src/utils"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Name:  "Nick",
		Email: "nick.test@gmail.com",
	},
	{
		Name:  "Byron",
		Email: "byron.test@gmail.com",
	},
}

var teams = []models.Team{
	{
		Name: "Cardinals",
		City: "Arizona",
	},
	{
		Name: "Falcons",
		City: "Atlanta",
	},
	{
		Name: "Ravens",
		City: "Baltimore",
	},
	{
		Name: "Bills",
		City: "Buffalo",
	},
	{
		Name: "Panthers",
		City: "Carolina",
	},
	{
		Name: "Bears",
		City: "Chicago",
	},
	{
		Name: "Bengals",
		City: "Cincinnati",
	},
	{
		Name: "Browns",
		City: "Cleveland",
	},
	{
		Name: "Cowboys",
		City: "Dallas",
	},
	{
		Name: "Broncos",
		City: "Denver",
	},
	{
		Name: "Lions",
		City: "Detroit",
	},
	{
		Name: "Packers",
		City: "Green Bay",
	},
	{
		Name: "Texans",
		City: "Houston",
	},
	{
		Name: "Colts",
		City: "Indianapolis",
	},
	{
		Name: "Jaguars",
		City: "Jacksonville",
	},
	{
		Name: "Chiefs",
		City: "Kansas City",
	},
	{
		Name: "Raiders",
		City: "Las Vegas",
	},
	{
		Name: "Chargers",
		City: "Los Angeles",
	},
	{
		Name: "Rams",
		City: "Los Angeles",
	},
	{
		Name: "Dolphins",
		City: "Miami",
	},
	{
		Name: "Vikings",
		City: "Minnesota",
	},
	{
		Name: "Patriots",
		City: "New England",
	},
	{
		Name: "Saints",
		City: "New Orleans",
	},
	{
		Name: "Giants",
		City: "New York",
	},
	{
		Name: "Jets",
		City: "New York",
	},
	{
		Name: "Eagles",
		City: "Philadelphia",
	},
	{
		Name: "Steelers",
		City: "Pittsburgh",
	},
	{
		Name: "49ers",
		City: "San Francisco",
	},
	{
		Name: "Seahawks",
		City: "Seattle",
	},
	{
		Name: "Buccaneers",
		City: "Tampa Bay",
	},
	{
		Name: "Titans",
		City: "Tennessee",
	},
	{
		Name: "Commanders",
		City: "Washington",
	},
}

var weeks = []models.Week{
	{
		Name:  "Week 1",
		Start: utils.ParseTime(time.RFC3339, "2022-07-07T16:41:46Z"),
		End:   utils.ParseTime(time.RFC3339, "2022-07-14T16:41:46Z"),
	},
	{
		Name:  "Week 2",
		Start: utils.ParseTime(time.RFC3339, "2022-07-15T16:41:46Z"),
		End:   utils.ParseTime(time.RFC3339, "2022-07-22T16:41:46Z"),
	},
}

var matchups = []models.Matchup{
	{
		HomeTeamID: 1,
		AwayTeamID: 2,
		WeekID:     1,
	},
	{
		HomeTeamID: 3,
		AwayTeamID: 4,
		WeekID:     1,
	},
	{
		HomeTeamID: 5,
		AwayTeamID: 6,
		WeekID:     1,
	},
	{
		HomeTeamID: 7,
		AwayTeamID: 8,
		WeekID:     1,
	},
	{
		HomeTeamID: 9,
		AwayTeamID: 10,
		WeekID:     1,
	},
}

var picks = []models.Pick{
	{
		MatchupID: 1,
		UserID:    1,
		WinnerID:  1,
	},
}

func LoadUsers(db *gorm.DB) {
	for i, _ := range users {
		err := db.Debug().Model(&models.User{}).Create(&users[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

func LoadTeams(db *gorm.DB) {
	for i, _ := range teams {
		err := db.Debug().Model(&models.Team{}).Create(&teams[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed teams table: %v", err)
		}
	}
}

func LoadWeeks(db *gorm.DB) {
	for i, _ := range weeks {
		err := db.Debug().Model(&models.Week{}).Create(&weeks[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed weeks table: %v", err)
		}
	}
}

func LoadMatchups(db *gorm.DB) {
	for i, _ := range matchups {
		err := db.Debug().Model(&models.Matchup{}).Create(&matchups[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed matchups table: %v", err)
		}
	}
}

func LoadPicks(db *gorm.DB) {
	for i, _ := range picks {
		err := db.Debug().Model(&models.Pick{}).Create(&picks[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed picks table: %v", err)
		}
	}
}

func LoadAll(db *gorm.DB) {
	LoadUsers(db)
	LoadTeams(db)
	LoadWeeks(db)
	LoadMatchups(db)
	LoadPicks(db)
}
