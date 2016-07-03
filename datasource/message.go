package datasource

import (
	"github.com/k0kubun/pp"
	m "github.com/timakin/airshooter/model"
)

func InsertMessage(message *m.Message) (*m.Message, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	dbConnection.Create(&message)
	var saved m.Message
	dbConnection.Preload("MessageSender").Preload("MessageRecipient").Last(&saved)

	return &saved, nil
}

func SelectMessages(threadId *int64) (messages *[]m.Message, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected []m.Message
	dbConnection.Preload("MessageSender").Preload("MessageRecipient").Where("thread_id = ?", &threadId).Find(&selected)

	return &selected, nil
}
