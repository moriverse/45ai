package models

import (
	"time"

	"gorm.io/gorm"
)

type Template struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Name          string         `gorm:"not null" json:"name"`
	Description   string         `json:"description"`
	CoverImageURL string         `json:"cover_image_url"`
	Price         float64        `gorm:"not null" json:"price"`
	StylePrompt   string         `gorm:"not null" json:"style_prompt"`
}
