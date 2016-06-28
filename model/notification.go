package model

// import (
// 	"github.com/jinzhu/gorm"
// )

type Notification struct {
	Id    *int64  `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	Title *string `json:"title"  validate:"required"`
	Text  *string `json:"text"   validate:"required"`
	// Status    *string `json:"status" validate:"required"`
	CreatedAt *int64 `json:"created_at"`
	UpdatedAt *int64 `json:"updated_at"`
	// DeliversAt *int64  `json:"delivers_at" validate:"required"`
	ExpiresAt *int64                `json:"expires_at"`
	Sender    NotificationSender    `json:"sender"`
	Recipient NotificationRecipient `json:"recipient"`
}

type Notifications []*Notification

func NewNotification(args map[string]interface{}) *Notification {
	return &Notification{
		Id:    args["id"].(*int64),
		Title: args["title"].(*string),
		Text:  args["text"].(*string),
		// Status:  args["status"].(*string),
		Sender:    args["sender"].(NotificationSender),
		Recipient: args["recipient"].(NotificationRecipient),
		CreatedAt: args["created_at"].(*int64),
		UpdatedAt: args["updated_at"].(*int64),
		// DeliversAt: args["delivers_at"].(*int64),
		ExpiresAt: args["expires_at"].(*int64),
	}
}
