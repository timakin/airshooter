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

func GetNotification() *m.Notification {

}

func GetNotificationCollection() []*m.Notification {

}
