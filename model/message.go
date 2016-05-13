package model

type Message struct {
	Id        *int64        `validate:"required"`
	Text      *string       `json:"text" validate:"required"`
	From      *Communicator `json:"from" validate:"required"`
	To        *Communicator `json:"to" validate:"required"`
	ThreadId  *int64        `json:"thread_id" validate:"required,len=10"`
	CreatedAt *int64        `validate:"required,len=10"`
}

type Messages []*Message

func NewMessage(args map[string]interface{}) *Message {
	return &Message{
		Text:      args["text"].(*string),
		From:      args["from"].(*Communicator),
		To:        args["to"].(*Communicator),
		ThreadId:  args["thread_id"].(*int64),
		CreatedAt: args["created_at"].(*int64),
	}
}
