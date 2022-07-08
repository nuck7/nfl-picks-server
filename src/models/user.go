package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique; not null" json:"email"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Picks     []Pick    `gorm:"foreign_key:UserID"`
}
