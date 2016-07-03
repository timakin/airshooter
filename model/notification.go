package model

// import (
// 	"github.com/jinzhu/gorm"
// )

type Notification struct {
	Id                    *int64                `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	Title                 *string               `json:"title"  validate:"required"`
	Text                  *string               `json:"text"   validate:"required"`
	Status                *string               `json:"status"`
	CreatedAt             *int64                `json:"created_at"`
	UpdatedAt             *int64                `json:"updated_at"`
	ExpiresAt             *int64                `json:"expires_at"`
	NotificationSender    NotificationSender    `json:"sender"`
	NotificationRecipient NotificationRecipient `json:"recipient"`
}

type Notifications []*Notification
