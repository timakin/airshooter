package service

import (
	db "github.com/timakin/airshooter/datasource"
	m "github.com/timakin/airshooter/model"
	"time"
)

func EnqueueNotification(req string) (result *m.Notification, err error) {
	req["created_at"] = time.Now().Unix()
	req["expiry"] = time.Now().Unix() + constant.NotificationExpiryDuration

	value := m.NewNotification(req)

	if result, err = db.InsertNotification(value); err != nil {
		return nil, err
	}

	return result, nil
}

func GetNotification(id *int64) (result *m.Notification, err error) {
	if result, err = db.SelectNotification(id); err != nil {
		return nil, err
	}
	return result, nil
}

func GetNotifications() (results []*m.Notification, err error) {
	if results, err = db.SelectNotifications(); err != nil {
		return nil, err
	}

	return results, nil
}
