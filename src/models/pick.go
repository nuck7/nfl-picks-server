package models

import "time"

type Pick struct {
	ID        uint      `gorm:"autoIncrement; unique"`
	UserID    uint      `gorm:"not null; primary_key"`
	MatchupID uint      `gorm:"not null; primary_key"`
	WinnerID  uint      `gorm:"not null; primary_key"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
