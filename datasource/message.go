package datasource

import (
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

func SelectMessages(params *map[string]interface{}) (messages *[]m.Message, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected []m.Message
	dbConnection.Preload("MessageSender").Preload("MessageRecipient").Where(&params).Find(&selected)

	return &selected, nil
}

// Select by recipientId and group by senderId
func SelectThreads(recipientId *int64) (messages *[]m.Message, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected []m.Message
	dbConnection.Preload("MessageSender").Preload("MessageRecipient").Where("message_recipients.recipient_id = ?", &recipientId).Group("message_senders.sender_id").Find(&selected)

	return &selected, nil
}
