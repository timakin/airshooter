package datasource

import (
	"github.com/pkg/errors"

	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
)

func InsertNotification(notification *m.Notification) (result *m.Notification, err error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}
	err = dbConnection.Insert(notification)
	if err != nil {
		return nil, errors.Wrap(err, constant.ErrDBInsertionFailed)
	}

	if err = dbConnection.SelectOne(result, "SELECT * FROM notifications ORDER BY created_at DESC LIMIT 1"); err != nil {
		return nil, errors.Wrap(err, constant.ErrDBSelectionFailed)
	}

	return result, nil
}

func SelectNotification(id *int64) (result *m.Notification, err error) {
	if dbConnection, err := GetDBInstance(); err != nil {
		return nil, err
	}

	if err = dbConnection.SelectOne(result, "select * from notifications where id=?", id); err != nil {
		return nil, errors.Wrap(err, constant.ErrDBSelectionFailed)
	}

	return result, nil
}

func SelectNotifications() ([]*m.Notification, error) {
	if dbConnection, err := GetDBInstance(); err != nil {
		return nil, err
	}

	var notifications []*m.Notification
	if _, err := dbConnection.Select(notifications, "select * from notifications order by created_at"); err != nil {
		return nil, errors.Wrap(err, constant.ErrDBSelectionFailed)
	}
	return notifications, nil
}
