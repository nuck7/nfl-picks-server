package models

import (
	"time"

	"gorm.io/gorm"
)

type Week struct {
	gorm.Model
	ID      uint      `gorm:"primaryKey"`
	Start   time.Time `gorm:"type:time" json:"start,omitempty"`
	Matches string    `json:"matches,omitempty"`
	Created int64     `gorm:"autoCreateTime"`
	Updated int64     `gorm:"autoUpdateTime"`
}
