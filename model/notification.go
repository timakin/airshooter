package model

import (
	db "github.com/timakin/airshooter/datasource"
)

type Notification struct {
	Id        *int64        `json:"id" db:"id, primarykey, autoincrement"`
	Title     *string       `json:"title" validate:"required" db:"title"`
	Text      *string       `json:"text" validate:"required" db:"text"`
	From      *Communicator `json:"from" validate:"required" db:"from"`
	To        *Communicator `json:"to" validate:"required" db:"to"`
	CreatedAt *int64        `json:"created_at" validate:"required,len=10" db:"created_at"`
	Expiry    *int64        `json:"expiry" validate:"required,len=10" db:"expiry"`
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

func PostNotification(args map[string]interface{}) (result *Notification, err error) {
	notification := NewNotification(args)
	if result, err = db.InsertNotification(notification); err != nil {
		return nil, err
	}

	return result, nil
}

func GetNotification(id *int64) (result *Notification, err error) {
	if result, err = db.SelectNotification(&id); err != nil {
		return nil, err
	}

	return result, nil
}

func GetNotifications() (result *Notification, err error) {
	if result, err = db.SelectNotifications(); err != nil {
		return nil, err
	}

	return result, nil
}
