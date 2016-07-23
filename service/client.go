package service

import (
	db "github.com/timakin/airshooter/datasource"
	m "github.com/timakin/airshooter/model"
	"time"
)

func GetClient(clientId *string) (client *m.Client, err error) {
	if client, err = db.SelectClient(clientId); err != nil {
		return nil, err
	}
	return client, nil
}

func RegisterClient() (client *m.Client, err error) {
	uid := 
	secret := 
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	if result, err := db.InsertClient(); err != nil {
		return nil, err
	}

	return result, nil
}

func InsertToken(token *m.AccessToken) (result *m.AccessToken, err error) {
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	token.CreatedAt = &createdAt
	token.UpdatedAt = &updatedAt

	if result, err = db.InsertToken(token); err != nil {
		return nil, err
	}

	return result, nil
}
