package db

import (
	"github.com/pkg/errors"

	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
)

func PostNotification(notification *m.Notification) error {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return err
	}
	err = dbConnection.Insert(notification)
	if err != nil {
		return errors.Wrap(err, constant.ErrDBInsertionFailed)
	}
	return nil
}

func GetNotification(id *int64) (*m.Notification, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var notification *m.Notification
	err = dbConnection.SelectOne(notification, "select * from notifications where id=?", id)
	if err != nil {
		return nil, errors.Wrap(err, constant.ErrDBSelectionFailed)
	}
	return notification, nil
}

func GetNotificationCollection() ([]*m.Notification, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var notifications []*m.Notification
	_, err = dbConnection.Select(notifications, "select * from notifications order by created_at")
	if err != nil {
		return nil, errors.Wrap(err, constant.ErrDBSelectionFailed)
	}
	return notifications, nil
}
