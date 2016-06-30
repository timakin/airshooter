package datasource

import (
	//	"github.com/pkg/errors"

	//	"github.com/timakin/airshooter/constant"
	"github.com/k0kubun/pp"
	m "github.com/timakin/airshooter/model"
)

func InsertNotification(notification *m.Notification) (*m.Notification, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	dbConnection.Create(&notification)
	var saved m.Notification
	dbConnection.Preload("NotificationSender").Preload("NotificationRecipient").Last(&saved)

	return &saved, nil
}

func SelectNotification(id *int64) (*m.Notification, error) {
	pp.Print(id)
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected m.Notification
	dbConnection.Preload("NotificationSender").Preload("NotificationRecipient").First(&selected, *id)

	return &selected, nil
}

func SelectNotifications() (notifications *[]m.Notification, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected []m.Notification
	dbConnection.Preload("NotificationSender").Preload("NotificationRecipient").Find(&selected)

	return &selected, nil
}

func PublishNotifications() (err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return err
	}

	dbConnection.Table("notifications").Where("status = waiting OR status = failed").Update("status", "published")

	return nil
}
