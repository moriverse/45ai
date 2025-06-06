package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskStatusPendingPayment TaskStatus = "PENDING_PAYMENT"
	TaskStatusProcessing     TaskStatus = "PROCESSING"
	TaskStatusCompleted      TaskStatus = "COMPLETED"
	TaskStatusFailed         TaskStatus = "FAILED"
)

type Task struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	UserID         uint           `gorm:"not null" json:"user_id"`
	User           User           `gorm:"foreignKey:UserID" json:"user"`
	TemplateID     uint           `gorm:"not null" json:"template_id"`
	Template       Template       `gorm:"foreignKey:TemplateID" json:"template"`
	Status         TaskStatus     `gorm:"not null" json:"status"`
	InputImageURL  string         `gorm:"not null" json:"input_image_url"`
	ResultImageURL string         `json:"result_image_url"`
	ErrorMessage   string         `json:"error_message"`
	WechatPrepayID string         `json:"wechat_prepay_id"`
}
