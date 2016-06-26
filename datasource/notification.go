package datasource

import (
	//	"github.com/pkg/errors"

	//	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
)

func InsertNotification(notification *m.Notification) (result *m.Notification, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	dbConnection.Create(&notification)
	dbConnection.First(result)

	return nil, nil
}

// func SelectNotification(id *int64) (result *m.Notification, err error) {
// 	dbConnection, err := GetDBInstance()
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return nil, nil
// }
//
// func SelectNotifications() (notifications []*m.Notification, err error) {
// 	dbConnection, err := GetDBInstance()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }
