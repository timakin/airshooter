package datasource

import (
	// "github.com/k0kubun/pp"
	m "github.com/timakin/airshooter/model"
)

func SelectClient(clientId *string, clientSecret *string) (client *m.Client, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected m.Client
	dbConnection.Where("uid = ? AND secret = ?", &clientId, &clientSecret).Find(&selected)

	return &selected, nil
}
