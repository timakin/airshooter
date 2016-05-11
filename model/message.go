package model

type Message struct {
	Text      *string `json: "text"`
	From      *Communicator
	To        *Communicator
	ThreadId  *int64
	CreatedAt *int64
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
