package service

import (
	db "github.com/timakin/airshooter/datasource"
	m "github.com/timakin/airshooter/model"
	"time"
)

func GetClient(clientId *string)

func InsertToken(message *m.Message) (result *m.Message, err error) {
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	message.CreatedAt = &createdAt
	message.UpdatedAt = &updatedAt

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
