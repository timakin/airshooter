package datasource

import (
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
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected m.Notification
	dbConnection.Preload("NotificationSender").Preload("NotificationRecipient").First(&selected, *id)

	return &selected, nil
}

func SelectNotifications(params *map[string]interface{}) (notifications *[]m.Notification, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var selected []m.Notification
	dbConnection.Preload("NotificationSender").Preload("NotificationRecipient").Where(&params).Find(&selected)

	return &selected, nil
}

func MarkNotifications(userId *int64) (results *[]m.Notification, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	condition := map[string]interface{}{"userId": &userId, "status": "unread"}
	dbConnection.Model(&results).Where(&condition).Update("status", "read")

	return results, nil
}
