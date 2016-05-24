package controller

import (
	m "github.com/timakin/airshooter/model"
)

type NotificationPostRequest struct {
	Title *string       `json:"title" validate:"required"`
	Text  *string       `json:"text" validate:"required"`
	From  *Communicator `json:"from" validate:"required"`
	To    *Communicator `json:"to" validate:"required"`
}

type NotificationGetRequest struct {
	Id *int64 `json:"id" validate:"required"`
}
