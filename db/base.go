package dbConnect

import (
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

var sharedInstance *sql.DB

func GetInstance() *sql.DB {
    var err error
    if sharedInstance == nil {
        sharedInstance, err = sql.Open("mysql", "root:@/my_database")
        if err != nil {
            panic(err.Error())
        }
    }
    return sharedInstance
}
