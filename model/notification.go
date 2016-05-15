package model

import (
	db "github.com/timakin/airshooter/datasource"
)

type Notification struct {
	Id        *int64        `json:"id"`
	Title     *string       `json:"title" validate:"required"`
	Text      *string       `json:"text" validate:"required"`
	From      *Communicator `json:"from" validate:"required"`
	To        *Communicator `json:"to" validate:"required"`
	CreatedAt *int64        `json:"created_at" validate:"required,len=10"`
	Expiry    *int64        `json:"expiry" validate:"required,len=10"`
}

type Notifications []*Notification

func NewNotification(args map[string]interface{}) *Notification {
	return &Notification{
		Title:     args["title"].(*string),
		Text:      args["text"].(*string),
		From:      args["from"].(*Communicator),
		To:        args["to"].(*Communicator),
		CreatedAt: args["created_at"].(*int64),
		Expiry:    args["expiry"].(*int64),
	}
}

func Post(notification *Notification) {
	if err := db.InsertNotification(notification)
}
