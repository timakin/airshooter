package model

type Message struct {
	Text             *string          `json:"text" validate:"required"`
	MessageSender    MessageSender    `json:"sender"`
	MessageRecipient MessageRecipient `json:"recipient"`
	CreatedAt        *int64           `json:"created_at"`
	UpdatedAt        *int64           `json:"updated_at"`
}

type Messages []*Message
