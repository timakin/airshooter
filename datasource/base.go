package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/timakin/airshooter/constant"
)

var sharedInstance *gorm.DB

// Singleton
func GetDBInstance() (*gorm.DB, error) {
	if sharedInstance == nil {
		db, err := gorm.Open("mysql", "root:@/airshooter?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			return nil, errors.Wrap(err, constant.ErrDBConnectionFailed)
		}
		db.LogMode(true)
		sharedInstance = db
	}
	return sharedInstance, nil
}
