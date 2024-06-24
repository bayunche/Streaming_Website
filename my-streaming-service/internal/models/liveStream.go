package models

import (
	"time"
)

type LiveStream struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	StartTime   time.Time
	EndTime     time.Time
	StreamKey   string `gorm:"unique;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
