package seed

import (
	"errors"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/nuck7/nfl-picks-server/src/models"
	"gorm.io/gorm"
)

func LoadUsers(db *gorm.DB) {
	for i, _ := range Users {
		err := db.Debug().Model(&models.User{}).Create(&Users[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

func LoadTeams(db *gorm.DB) {
	for i, _ := range Teams {
		err := db.Debug().Model(&models.Team{}).Create(&Teams[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed teams table: %v", err)
		}
	}
}

func LoadWeeks(db *gorm.DB) {
	for i, _ := range Weeks {
		err := db.Debug().Model(&models.Week{}).Create(&Weeks[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed weeks table: %v", err)
		}
	}
}

func LoadMatchups(db *gorm.DB) {
	for i, _ := range Matchups {
		err := db.Debug().Model(&models.Matchup{}).Create(&Matchups[i]).Error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number != 1062 {
			log.Fatalf("cannot seed matchups table: %v", err)
		}
	}
}

func LoadPicks(db *gorm.DB) {
	for i, _ := range Picks {
		err := db.Debug().Model(&models.Pick{}).Create(&Picks[i]).Error
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
