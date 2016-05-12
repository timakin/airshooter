package model

type Message struct {
	Text      *string       `json: "text"`
	From      *Communicator `json: "from"`
	To        *Communicator `json: "to"`
	ThreadId  *int64        `json: "thread_id"`
	CreatedAt *int64        `json: "created_at"`
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
