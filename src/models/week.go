package models

import (
	"time"
)

type Week struct {
	ID        uint      `gorm:"primary_key"`
	Start     time.Time `gorm:"type:DATETIME; unique; not null"`
	End       time.Time `gorm:"type:DATETIME; unique; not null"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Matchups  []Matchup `gorm:"foreignKey:WeekID"`
}
