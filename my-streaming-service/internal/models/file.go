package models

import (
	"time"
)

type File struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	FileType  string `gorm:"not null"`
	FilePath  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
