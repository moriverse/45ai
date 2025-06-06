package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	OpenID    string         `gorm:"uniqueIndex;not null" json:"openid"`
	Nickname  string         `json:"nickname"`
	AvatarURL string         `json:"avatar_url"`
}
