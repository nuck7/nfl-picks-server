package models

import "time"

type Team struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"unique; not null"`
	City      string    `gorm:"unqiue; not null"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Winners   []Pick    `gorm:"foreignKey:WinnerID"`
	HomeTeams []Matchup `gorm:"foreignKey:HomeTeamID"`
	AwayTeams []Matchup `gorm:"foreignKey:AwayTeamID"`
}
