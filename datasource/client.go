package datasource

import (
	// "github.com/k0kubun/pp"
	m "github.com/timakin/airshooter/model"
)

func SelectClient(clientId *string) (client *m.Client, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected m.Client
	dbConnection.Where("uid = ? AND secret = ?", &clientId).Find(&selected)

	return &selected, nil
}

func InsertToken(token *m.AccessToken) (*m.AccessToken, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	dbConnection.Create(&token)
	var saved m.AccessToken
	dbConnection.Last(&saved)

	return &saved, nil
}
