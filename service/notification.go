package service

import (
	constant "github.com/timakin/airshooter/constant"
	db "github.com/timakin/airshooter/datasource"
	m "github.com/timakin/airshooter/model"
	"time"
)

func EnqueueNotification(notification *m.Notification) (result *m.Notification, err error) {
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()
	expiresAt := time.Now().Unix() + constant.NotificationExpiryDuration

	sender := m.NotificationSender{
		SenderId:   notification.NotificationSender.SenderId,
		SenderType: notification.NotificationSender.SenderType,
		CreatedAt:  &createdAt,
		UpdatedAt:  &updatedAt,
	}
	recipient := m.NotificationRecipient{
		RecipientId:   notification.NotificationRecipient.RecipientId,
		RecipientType: notification.NotificationRecipient.RecipientType,
		CreatedAt:     &createdAt,
		UpdatedAt:     &updatedAt,
	}

	notification.ExpiresAt = &expiresAt
	notification.CreatedAt = &createdAt
	notification.UpdatedAt = &updatedAt
	notification.NotificationSender = sender
	notification.NotificationRecipient = recipient

	if result, err = db.InsertNotification(notification); err != nil {
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

func GetAllNotifications(userId *int64) (results *[]m.Notification, err error) {
	if results, err = db.SelectNotifications(); err != nil {
		return nil, err
	}

	return results, nil
}

func GetUnreadNotifications(userId *int64) (results *[]m.Notification, err error) {
	if results, err = db.SelectNotifications(); err != nil {
		return nil, err
	}

	return results, nil
}

func MarkNotificationsAsRead(userId *int64) (err error) {
	if updated, err = db.PublishNotifications(); err != nil {
		return nil, err
	}

	return nil
}
