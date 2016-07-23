package service

import (
	db "github.com/timakin/airshooter/datasource"
	m "github.com/timakin/airshooter/model"
	sr "github.com/tuvistavie/securerandom"
	"time"
)

func GetClient(clientId *string) (client *m.Client, err error) {
	if client, err = db.SelectClient(clientId); err != nil {
		return nil, err
	}
	return client, nil
}

func RegisterClient() (result *m.Client, err error) {
	uid, _ := sr.Hex(10)
	secret, _ := sr.Hex(10)
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	client := m.Client{
		UID:       &uid,
		Secret:    &secret,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}

	if result, err = db.InsertClient(&client); err != nil {
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
