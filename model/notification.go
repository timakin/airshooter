package model

import (
    "time"
)

// from, to = id, type
// CreatedAt, Expiry = unixtimestamp
type Notification struct {
	Title     *string `json: "title"`
	Text      *string `json: "text"`
	From      *Communicator
	To        *Communicator
	CreatedAt *int64 `json: "created_at"`
	Expiry    *int64 `json: "expiry"`
}

type Notifications []*Notification

func newNotification(title string, text string, from Communicator, to Communicator, createdAt int64, expiry int64) *Notification {
    now := time.Now().UnixNano()

    return &Notification {
        Title: &title,
        Text: &text,
        From: &from,
        To: &to,
        CreatedAt: &now,
        Expiry: &now,
    }
}
