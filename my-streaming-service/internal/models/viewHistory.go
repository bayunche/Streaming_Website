package models

import (
	"time"
)

type ViewHistory struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	VideoID   uint      `gorm:"not null"`
	WatchedAt time.Time `gorm:"default:current_timestamp"`
}
