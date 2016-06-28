package model

type NotificationRecipient struct {
	Id             *int64  `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	NotificationId *int64  `json:"notification_id"`
	RecipientId    *int64  `json:"recipient_id" validate:"required"`
	Sector         *string `json:"sector" validate:"required"`
	CreatedAt      *int64  `json:"created_at"`
	UpdatedAt      *int64  `json:"updated_at"`
}
