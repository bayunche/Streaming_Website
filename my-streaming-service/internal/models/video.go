package models

import (
	"time"
)

type Video struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"not null"`
	Title         string `gorm:"not null"`
	Description   string
	FilePath      string `gorm:"not null"`
	ThumbnailPath string
	IsLiveRecord  bool `gorm:"default:false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
