package model

type Message struct {
	Text             *string          `json:"text" validate:"required"`
	ThreadId         *int64           `json:"thread_id" validate:"required"`
	MessageSender    MessageSender    `json:"sender"`
	MessageRecipient MessageRecipient `json:"recipient"`
}

type Messages []*Message
