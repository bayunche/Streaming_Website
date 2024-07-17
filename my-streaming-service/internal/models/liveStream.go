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

type LiveRoom struct {
	ID          uint   `gorm:"primaryKey"`
	RoomName    string `gorm:"unique;not null"`
	Description string `gorm:"unique;not null"`
	UserID      string `gorm:"not null"`
	RoomUrl     string `gorm:"unique;not null"`
	RoomId      string `gorm:"unique;not null"`
	createdAt   time.Time
}
