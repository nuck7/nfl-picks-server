package models

import (
	"time"
)

type Matchup struct {
	ID         uint      `gorm:"primary_key"`
	HomeTeamID uint      `gorm:"not null"`
	AwayTeamID uint      `gorm:"not null"`
	WeekID     uint      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Picks      []Pick    `gorm:"foreignKey:MatchupID"`
}
