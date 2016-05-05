package model

type Message struct {
	Text      *string `json: "text"`
	From      *Communicator
	To        *Communicator
	ThreadId  *int
	CreatedAt *int
}
