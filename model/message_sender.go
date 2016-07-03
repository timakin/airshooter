package model

type MessageSender struct {
	Id         *int64  `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	MessageId  *int64  `json:"message_id"`
	SenderId   *int64  `json:"sender_id" validate:"required`
	SenderType *string `json:"type" validate:"required`
	CreatedAt  *int64  `json:"created_at"`
	UpdatedAt  *int64  `json:"updated_at"`
}
