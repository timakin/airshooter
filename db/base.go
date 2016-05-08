package db

import (
    "github.com/timakin/airshooter/constant"

    "database/sql"
    "gopkg.in/gorp.v1"
    "github.com/pkg/errors"

    m "github.com/timakin/airshooter/model"
    _ "github.com/go-sql-driver/mysql"
)

var sharedInstance *gorp.DbMap

// Singleton
func GetDBInstance() (*gorp.DbMap, error) {
    if sharedInstance == nil {
        connection, err := sql.Open("mysql", "root:@/my_database")
        if err != nil {
            return nil, errors.Wrap(err, constant.ErrDBConnectionFailed)
        }
        sharedInstance = &gorp.DbMap{Db: connection, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
        sharedInstance.AddTableWithName(m.Notification{}, "notifications").SetKeys(true, "Id")
        sharedInstance.AddTableWithName(m.Message{}, "messages").SetKeys(true, "Id")
    }
    return sharedInstance, nil
}

