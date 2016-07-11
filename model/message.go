package model

type Message struct {
	Id               *int64           `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	Text             *string          `json:"text" validate:"required"`
	CreatedAt        *int64           `json:"created_at"`
	UpdatedAt        *int64           `json:"updated_at"`
	MessageSender    MessageSender    `json:"sender"`
	MessageRecipient MessageRecipient `json:"recipient"`
}

type Messages []*Message
