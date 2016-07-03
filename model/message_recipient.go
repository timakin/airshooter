package model

type MessageRecipient struct {
	Id            *int64  `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	MessageId     *int64  `json:"message_id"`
	RecipientId   *int64  `json:"recipient_id" validate:"required"`
	RecipientType *string `json:"type" validate:"required"`
	CreatedAt     *int64  `json:"created_at"`
	UpdatedAt     *int64  `json:"updated_at"`
}
