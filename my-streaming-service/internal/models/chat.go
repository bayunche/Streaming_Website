package models

import (
	"time"
)

type ChatMessage struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    StreamID  uint      `gorm:"not null"`
    Message   string    `gorm:"not null"`
    CreatedAt time.Time
}
