package service

import (
	m "github.com/timakin/airshooter/model"
	"time"
)

func EnqueueNotification(req string) (result *m.Notification, err error) {
	req["created_at"] = time.Now().Unix()
	req["expiry"] = time.Now().Unix() + constant.NotificationExpiryDuration

	if result, err = m.PostNotification(req); err != nil {
		return nil, err
	}

	return result, nil
}

func GetNotification(id *int64) (result *m.Notification, err error) {
	if result, err = m.GetNotification(id); err != nil {
		return nil, err
	}
	return result, nil
}

func GetNotifications() (results []*m.Notification, err error) {
	if results, err = m.GetNotifications(); err != nil {
		return nil, err
	}

	return results, nil
}
