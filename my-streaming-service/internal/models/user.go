package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Role      string `gorm:"default:user"`
	UserID    string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Avatar    string
	StreamUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
