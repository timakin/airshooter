package service

import (
	db "github.com/timakin/airshooter/datasource"
	m "github.com/timakin/airshooter/model"
	"time"
)

func GetClient(clientId *string) (client *m.Client, err error) {
	if results, err = db.SelectClient(clientId); err != nil {
		return nil, err
	}
	return results, ni
}

func InsertToken(token *m.AccessToken) (result *m.AccessToken, err error) {
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	message.CreatedAt = &createdAt
	message.UpdatedAt = &updatedAt

	if result, err = db.InsertToken(token); err != nil {
		return nil, err
	}

	return result, nil
}
