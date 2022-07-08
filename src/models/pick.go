package models

import "time"

type Pick struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `gorm:"not null" json:"userID"`
	MatchupID uint      `gorm:"not null" json:"matchupID"`
	WinnerID  uint      `gorm:"not null" json:"winnerID"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
