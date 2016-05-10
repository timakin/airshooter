package datasource

import (
	"github.com/pkg/errors"

	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
)

func PostMessage(message *m.Message) error {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return err
	}
	err = dbConnection.Insert(message)
	if err != nil {
		return errors.Wrap(err, constant.ErrDBInsertionFailed)
	}
	return nil
}

func GetMessage(id *int64) (*m.Message, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var message *m.Message
	err = dbConnection.SelectOne(message, "select * from messages where id=?", id)
	if err != nil {
		return nil, errors.Wrap(err, constant.ErrDBSelectionFailed)
	}
	return message, nil
}

func GetMessages() ([]*m.Message, error) {
	dbConnection, err := GetDBInstance()
	if err != nil {
		return nil, err
	}

	var messages []*m.Message
	_, err = dbConnection.Select(messages, "select * from messages order by created_at")
	if err != nil {
		return nil, errors.Wrap(err, constant.ErrDBSelectionFailed)
	}
	return messages, nil
}
