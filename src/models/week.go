package models

import (
	"time"
)

type Week struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Start     time.Time `gorm:"type:DATETIME; not null"`
	End       time.Time `gorm:"type:DATETIME; not null"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Matchups  []Matchup `gorm:"foreignKey:WeekID"`
}
