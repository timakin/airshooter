package model

import (
	"time"
)

type Message struct {
	Text      *string `json: "text"`
	From      *Communicator
	To        *Communicator
	ThreadId  *int64
	CreatedAt *int64
}

type Messages []*Message

func newMessage(text string, from Communicator, to Communicator, threadId int64, createdAt int64) *Message {
	now := time.Now().UnixNano()

	return &Message{
		Text:      &text,
		From:      &from,
		To:        &to,
		ThreadId:  &threadId,
		CreatedAt: &now,
	}
}
