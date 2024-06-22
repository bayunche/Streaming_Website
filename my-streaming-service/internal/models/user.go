package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 定义一个User结构体

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
	UserName  string    ` gorm:"not null"`
	UserId    string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Phone     string    `gorm:"not null"`
	Address   string    `gorm:"not null"`
	Status    int       `gorm:"not null"`
	Level     int       `gorm:"not null"`
	Role      int       `gorm:"not null"`
}
