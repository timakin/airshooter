package datasource

import (
	//	"github.com/pkg/errors"

	//	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
)

func InsertNotification(notification *m.Notification) (*m.Notification, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	dbConnection.Create(&notification)
	var saved m.Notification
	dbConnection.Order("id desc").Limit(1).Find(&saved)

	return &saved, nil
}

func SelectNotification(id *int64) (*m.Notification, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected m.Notification
	var sender m.NotificationSender
	dbConnection.Find(&selected, id).Related(&sender)

	return &selected, nil
}

func SelectNotifications() (notifications *[]m.Notification, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected []m.Notification
	var sender m.NotificationSender
	dbConnection.Model(&selected).Related(&sender, "Sender")

	return &selected, nil
}
