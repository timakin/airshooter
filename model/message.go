package model

type Message struct {
	Text      *string `json: "text"`
	From      *string
	To        *string
	ThreadId  *int
	CreatedAt *int
}
