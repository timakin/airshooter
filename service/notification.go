package service

import (
	m "github.com/timakin/airshooter/model"
	"time"
)

func EnqueueNotification(req string) error {
	req["created_at"] = time.Now().Unix()
	req["expiry"] = time.Now().Unix() + constant.NotificationExpiryDuration

	if err := m.PostNotification(req); err != nil {
		return err
	}

	return nil
}

func GetNotification(id int64) (*m.Notification, error) {
	return nil, nil
}

func GetNotifications(ids []int64) ([]*m.Notification, error) {
	// find notifications by ids
	return nil, nil
}
