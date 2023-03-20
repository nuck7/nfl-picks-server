package models

import (
	"time"
)

type Matchup struct {
	ID         uint      `gorm:"autoIncrement; primary_key"`
	HomeTeamID uint      `gorm:"not null; index:idx_schedule, unique"`
	AwayTeamID uint      `gorm:"not null; index:idx_schedule, unique"`
	WeekID     uint      `gorm:"not null; index:idx_schedule, unique"`
	CreatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Picks      []Pick    `gorm:"foreignKey:MatchupID;References:ID"`
}
