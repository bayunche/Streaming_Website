package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 定义一个流媒体播放资源结构体
type Media struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	DeletedAt     time.Time `gorm:"index"`
	MediaId       string    `gorm:"not null"`
	MediaName     string    ` gorm:"not null"`
	MediaUrl      string    `gorm:"not null"`
	MediaType     int       `gorm:"not null"`
	MediaSize     int       `gorm:"not null"`
	MediaCover    string    `gorm:"not null"`
	MediaDesc     string    `gorm:"not null"`
	MediaDuration int       `gorm:"not null"`
	MediaStatus   int       `gorm:"not null"`
	MediaUserId   string    `gorm:"not null"`
}
