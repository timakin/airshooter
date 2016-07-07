package service

import (
	db "github.com/timakin/airshooter/datasource"
	m "github.com/timakin/airshooter/model"
	"time"
)

func InsertMessage(message *m.Message) (result *m.Message, err error) {
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	sender := m.MessageSender{
		SenderId:   message.MessageSender.SenderId,
		SenderType: message.MessageSender.SenderType,
		CreatedAt:  &createdAt,
		UpdatedAt:  &updatedAt,
	}
	recipient := m.MessageRecipient{
		RecipientId:   message.MessageRecipient.RecipientId,
		RecipientType: message.MessageRecipient.RecipientType,
		CreatedAt:     &createdAt,
		UpdatedAt:     &updatedAt,
	}

	message.CreatedAt = &createdAt
	message.UpdatedAt = &updatedAt
	message.MessageSender = sender
	message.MessageRecipient = recipient

	if result, err = db.InsertMessage(message); err != nil {
		return nil, err
	}

	return result, nil
}

func GetMessages(params *map[string]interface{}) (results *[]m.Message, err error) {
	if results, err = db.SelectMessages(params); err != nil {
		return nil, err
	}

	return results, nil
}

func GetThreads(recipientId *int64) (results *[]m.Message, err error) {
	if results, err = db.SelectThreads(recipientId); err != nil {
		return nil, err
	}
	return results, nil
}
