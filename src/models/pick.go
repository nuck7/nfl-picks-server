package models

import "time"

type Pick struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `gorm:"not null"`
	MatchupID uint      `gorm:"not null"`
	WinnerID  uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
